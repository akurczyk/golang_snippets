package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	//
	// strings -- basic operations on strings
	// Based on: https://gobyexample.com/string-functions
	//

	fmt.Println(strings.Contains("abcdef", "abc"))
	fmt.Println(strings.Count("aaa", "a"))
	fmt.Println(strings.HasPrefix("abcdef", "abc"))
	fmt.Println(strings.HasSuffix("abcdef", "def"))
	fmt.Println(strings.Index("abcdef", "d"))
	fmt.Println(strings.Join([]string{"abc", "def"}, ","))
	fmt.Println(strings.Repeat("a", 10))
	fmt.Println(strings.Replace("abcdefabc", "abc", "xyz", 1)) // Only first
	fmt.Println(strings.Replace("abcdefabc", "abc", "xyz", 2)) // First two
	fmt.Println(strings.Replace("abcdefabc", "abc", "xyz", -1)) // All
	fmt.Println(strings.Split("abc,def", ","))
	fmt.Println(strings.ToLower("AbcDef"))
	fmt.Println(strings.ToUpper("AbcDef"))
	fmt.Println()



	//
	// strconv -- numbers parsing
	// Source: https://gobyexample.com/number-parsing
	//

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	fmt.Println()



	//
	// regexp -- regular expressions
	// Source: https://gobyexample.com/regular-expressions
	//

	// Short match
	match, _ := regexp.MatchString("aaa[0-9]+zzz", "aaa123123123zzz")
	fmt.Println(match)

	r, _ := regexp.Compile("aaa[0-9]+zzz")
	// r := regexp.MustCompile("aaa[0-9]+zzz") // Panics instead of returning errors

	// Match and find
	fmt.Println("MS:", r.MatchString("aaa999zzz"))
	fmt.Println("M:", r.Match([]byte("aaa999zzz")))
	fmt.Println("FS:", r.FindString("aaa111zzz zzzaaa aaa222zzz"))
	fmt.Println("FSI:", r.FindStringIndex("aaa111zzz zzzaaa aaa222zzz"))
	fmt.Println("FSS:", r.FindStringSubmatch("aaa111zzz zzzaaa aaa222zzz"))
	fmt.Println("FSSI:", r.FindStringSubmatchIndex("aaa111zzz zzzaaa aaa222zzz"))
	fmt.Println("FAS(1):", r.FindAllString("aaa111zzz zzzaaa aaa222zzz", 1))
	fmt.Println("FAS(2):", r.FindAllString("aaa111zzz zzzaaa aaa222zzz", 2))
	fmt.Println("FAS(-1):", r.FindAllString("aaa111zzz zzzaaa aaa222zzz", -1))
	fmt.Println()

	// Replace
	fmt.Println(r.ReplaceAllString("aaa123zzz aaa999999999zzz", "<here was our regexp>"))
	fmt.Println(r.ReplaceAllStringFunc("aaa123zzz aaa999999999zzz", strings.ToUpper))
	fmt.Println()



	//
	// time -- time formatting and parsing
	// Source: https://gobyexample.com/time-formatting-parsing
	//

	t1 := time.Now()
	fmt.Println(t1)

	t2, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println(t2)

	// The format string must be based on this date: 2006-01-02T15:04:05Z07:00
	t3, _ := time.Parse("3:04 PM", "8:41 PM")
	fmt.Println(t3)
}
