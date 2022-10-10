package ScrapeAndParse

func GoogleScrape(searchTerm string, countryCode string, languageCode string, proxyString interface{}, pages int, count int, backoff int) ([]searchResult, error) {
	results := []searchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}
	for _, page := range googlePages {
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := googleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			results = append(results, result)
		}
		//time.Sleep(time.Second * time.Duration(backoff))
	}
	return results, nil
}
