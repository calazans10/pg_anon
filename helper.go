package main

import "regexp"

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Gsub(subject, patern string) string {
	re := regexp.MustCompile(patern)
	match := re.FindStringSubmatch(subject)
	return match[1]
}

func Replace(src, pattern, repl string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(src, repl)
}
