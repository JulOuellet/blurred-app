package searchbars

type SearchResultItem struct {
	Type  string
	Name  string
	URL   string
	Extra string
}

type resultGroup struct {
	Label   string
	Results []SearchResultItem
}

func groupResults(results []SearchResultItem) []resultGroup {
	order := []string{"sport", "championship", "event"}
	labels := map[string]string{
		"sport":        "Sports",
		"championship": "Championships",
		"event":        "Events",
	}

	grouped := make(map[string][]SearchResultItem)
	for _, r := range results {
		grouped[r.Type] = append(grouped[r.Type], r)
	}

	var groups []resultGroup
	for _, t := range order {
		if items, ok := grouped[t]; ok {
			groups = append(groups, resultGroup{
				Label:   labels[t],
				Results: items,
			})
		}
	}
	return groups
}
