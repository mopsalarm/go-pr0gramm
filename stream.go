package pr0gramm

type Consumer func(Item) (bool, error)

func Stream(req ItemsRequest, consume Consumer) error {
	for {
		items, err := GetItems(req)
		if err != nil {
			return err
		}

		for _, item := range items.Items {
			if cont, err := consume(item); err != nil {
				return err
			} else if !cont {
				return nil
			}
		}

		if len(items.Items) == 0 || items.AtEnd {
			return nil
		}

		req.Older = items.Items[len(items.Items)-1].Id
	}
}

func ConsumeIf(predicate func(Item) bool, consumer func(Item) error) Consumer {
	return func(item Item) (bool, error) {
		if predicate(item) {
			return true, consumer(item)
		} else {
			return false, nil
		}
	}
}
