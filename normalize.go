package normalize

import (
 "bytes"
 "unicode"
 "unicode/utf8"
)

func lowercase(runes []rune) []rune {
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
	return runes
}

func upperfirst(runes []rune) []rune {
	runes[0] = unicode.ToUpper(runes[0])
	for i, r := range runes[1:] {
		runes[i] = unicode.ToLower(r)
	}
	return runes
}

func isException(runes []rune) bool {
	if len(runes) < 3 {
		return true
	}
	switch string(runes) {
		case `ABC`, `USA`, `USSR`, `YMCA`, `RAF`, `USAF`, `USCG`, `USMC`, `USN`:
			return true
	}
	return false
}

func Text(b []byte) []byte {
	
	// Convert to slice of runes
	n := len(b)
	if n == 0 {
		return b
	}
	para := make([][][]rune, 0, 1)
	sent := make([][]rune, 0, 1)
	word := make([]rune, 0, 1)
	var r rune
	var i, w int
    for i=0; i<n; i+=w {
        r, w = utf8.DecodeRune(b[i:])

		switch r {
			// New word
			case ' ':
				if len(word) == 0 {
					continue
				}
				sent = append(sent, word)
				word = make([]rune, 0, 1)
				continue
			// New line
			case 10, 13:
				if len(sent) == 0 {
					continue
				}
				if len(word) > 0 {
					sent = append(sent, word)
					word = make([]rune, 0, 1)
				}
				para = append(para, sent)
				sent = make([][]rune, 0, 1)
				continue
			// Normalize aprostrophes
			case '‘', '’': r = 39
			case '“', '”': r = '"'
		}
		
		// Add rune
		word = append(word, r)
	}
	if len(word) > 0 {
		sent = append(sent, word)
	}
	if len(sent) > 0 {
		para = append(para, sent)
	}
	
	// Write it back
	buf := bytes.NewBuffer(make([]byte, 0, 20))
	var i2 int
	for i, sent = range para {
		if i > 0 {
			buf.WriteByte(10)
			if len(sent) > 10 {
				buf.WriteByte(10)
			}
		}
		for i2, word = range sent {
			if i2 > 0 {
				buf.WriteByte(' ')
			}
			for _, r = range word {
				buf.WriteRune(r)
			}
		}
	}
	return buf.Bytes()

	/*
	numpara := len(para)
	
	var ok, allcaps, firstcap, othercap, anyletter, puncbefore bool
	var last, i2, on int
	for i=0; i<numpara; i++ {
		sent = para[i]
		
		puncbefore = true
		// Loop though each word and correct casing
		for on, word = range sent {
		
			// First check for lost punctuation in front of the word
			ok = true
			for i2, r = range word {
				if unicode.IsLetter(r) || unicode.IsNumber(r) {
					break
				}
				switch r {
					case ',', '.', ':', ';', '!':
						ok = false
				}
			}
			if !ok {
				if on > 0 && !puncbefore && unicode.IsLetter(r) && (unicode.IsUpper(r) || unicode.IsTitle(r)) {
					sent[on-1] = append(sent[on-1], word[0:i2]...)
					puncbefore = true
				}
				sent[on] = word[i2:]
				word = word[i2:]
			}
		
			// Calculate the casing type of the word
			r = word[0]
			if unicode.IsLetter(r) {
				anyletter = true
				if unicode.IsUpper(r) || unicode.IsTitle(r) {
					allcaps = true
					firstcap = true
					othercap = true
				} else {
					allcaps = false
					firstcap = false
					othercap = false
				}
			} else {
				anyletter = false
				allcaps = false
				firstcap = false
				othercap = false
			}
			for _, r = range word {
				if unicode.IsLetter(r) {
					anyletter = true
					if unicode.IsUpper(r) || unicode.IsTitle(r) {
						othercap = true
					} else {
						allcaps = false
					}
				}
			}
			if anyletter {
				if puncbefore {
					if !firstcap {
						upperfirst(word)
					} else {
						if allcaps {
							if !isException(word) {
								upperfirst(word)
							}
						} else {
							if othercap {
								upperfirst(word)
							}
						}
					}
				} else {
					if othercap {
						if allcaps {
							if !isException(word) {
								lowercase(word)
							}
						} else {
							lowercase(word)
						}
					}
				}
			}
			
			puncbefore = false
			for i2=len(word)-1; i2>=0; i2-- {
				r = r
				if unicode.IsPunct(r) {
					switch r {
						case '.', '?', '!':
							puncbefore = true
							break
					}
				} else {
					break
				}
			}
			
		}
		
		// Ensure last character is suitable
		word = sent[len(sent) - 1]
		last = len(word) - 1
		ok = false
		switch word[last] {
			case '.', '!', '?', ')', '"', 39: ok = true
		}
		if !ok {
			for i2=last-1; i2>=0; i2-- {
				r = word[i2]
				switch r {
					case '.', '!', '?', ')', '"', 39:
						sent[len(sent) - 1] = sent[len(sent) - 1][0:i2+1]
						break
				}
				if unicode.IsLetter(r) || unicode.IsNumber(r) {
					word[i2+1] = '.'
					sent[len(sent) - 1] = sent[len(sent) - 1][0:i2+2]
					break
				}
			}
		}
	}
	
	// Write it back
	buf := bytes.NewBuffer(make([]byte, 0, 20))
	for i, sent = range para {
		if i > 0 {
			buf.WriteByte(10)
			if len(sent) > 10 {
				buf.WriteByte(10)
			}
		}
		for i2, word = range sent {
			if i2 > 0 {
				buf.WriteByte(' ')
			}
			for _, r = range word {
				buf.WriteRune(r)
			}
		}
	}
	return buf.Bytes()
	*/
}
