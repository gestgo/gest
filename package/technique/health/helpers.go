package health

func Up(key string, details map[string]any) IndicatorResult {
	return IndicatorResult{Key: key, Up: true, Details: details}
}

func Down(key string, err error, details map[string]any) IndicatorResult {
	r := IndicatorResult{Key: key, Up: false, Details: details}
	if err != nil {
		r.Error = err.Error()
	}
	return r
}
