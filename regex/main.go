package main

import (
	"fmt"
	"regexp"
)

/*
Regex = a pattern to find text

Example:
Pattern: go
Text: I like go
Result: ✅ found

# Literals → actual characters
Rules → what is allowed, how many times, position, etc.
*/

func main() {
	pattern := "go"

	text := "i like to go party"

	matched, _ := regexp.MatchString(pattern, text)

	if matched {
		fmt.Println("Found")
	}

}

/*

Simplest Regex (Literals)
cat

Matches:
✅ cat
❌ cats
❌ dog
👉 Exact match only.

Character Classes [ ]
Match one of many characters
[abc]

Matches:
a OR b OR c

Example:
gr[ae]y
Matches:
gray
grey



# Ranges
[a-z]   // lowercase letters
[A-Z]   // uppercase letters
[0-9]   // digits

Example:
[0-9][0-9]


Matches:
12
99
❌ 9




Quantifiers (HOW MANY times)
These are very important.
Symbol		Meaning
*			0 or more
+			1 or more
?			0 or 1
{n}			exactly n
{n,m}		between n and m

Examples
a*

Matches:
"" (empty)
a
aaaa
a+


Matches:
a
aaaa
❌ empty

[0-9]{10}
Matches:
9876543210
👉 Indian mobile number basic match



Special Shortcuts (VERY COMMON)
Shortcut			Meaning
\d					digit [0-9]
\w					word char [a-zA-Z0-9_]
\s					space / tab
.					any character


\d{4} : Digits of 4 times

Matches:
2025


Anchors (POSITION matters)
Anchor	Meaning
^	start of string
$	end of string


^\d{10}$
Matches:

9876543210
❌ Does NOT match:
abc9876543210xyz

👉 Very important for validation


Groups ( )

Capture parts of a match
(\d{2})-(\d{2})-(\d{4})

Matches:
15-12-2025

Groups:
15
12
2025

Used heavily in:
Parsing
Replacing
Code analysis


OR condition |
cat|dog

Matches:
cat
dog


^(yes|no)$



############################################

Escaping special characters \
Some characters have special meaning:
. + * ? ^ $ ( ) [ ] { } | \

To match them literally:
\.
\$
\(


Example:
\$
Matches $

############################################

############################################

Real-world examples (IMPORTANT)

✔ Email (basic)

^[a-zA-Z0-9._]+@[a-zA-Z]+\.[a-zA-Z]{2,}$


IP Address (basic)
^(\d{1,3}\.){3}\d{1,3}$


SQL Injection detection (like your Go analyzer)
(?i)(select|insert|update|delete|drop)\s
(?i) → case-insensitive

✔ Slightly Improved Version (Still Basic)
(?i)\b(select|insert|update|delete|drop)\b\s+

Improvements:
\b → word boundaries
\s+ → one or more spaces
Fewer false positives

#################################################


################################################
How to BUILD regex for YOUR needs (STEP METHOD)
🧠 Step-by-step approach

Write sample text
User ID: EMP-1023

Mark variable parts
EMP-[numbers]

Translate to regex
EMP-\d+


Add anchors (if validation)
^EMP-\d+$

##############################################


Regex in Go (your context)
re := regexp.MustCompile(`^\d{10}$`)
fmt.Println(re.MatchString("9876543210")) // true

##############################################

*/
