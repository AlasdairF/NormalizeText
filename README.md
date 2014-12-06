##NormalizeText

This package normalizes UTF8 text to make it look more 'pretty'. Specifically it's meant to clean up text that's come out of OCR, to make it at least partially presentable and minimize or hide mistakes.

##Features

* UTF8 compliant
* Supports all languages that use common punctuation `.,!?:;`
* Fast: does not use any regular expressions
* Normalizes all casing
* Normalizes word and line spacing
* Repairs broken lines
* Preserves uppercase on first letter of a word if only the first letter is capitalized
* Preserves ALLCAPS for roman numerals
* Preserves ALLCAPS for words of less than 3 characters (e.g. FL, NY, UK, etc.)
* Preserves ALLCAPS for the following exceptions: `ABC`, `USA`, `USSR`, `YMCA`, `RAF`, `USAF`, `USCG`, `USMC`, `USN`
* Corrects uppercase for words following punctuation
* Corrects uppercase for first word in sentence
* Converts fancy aprostrophes `‘’“”` to common `'"`
* Removes out-of-place punctuation
* Normalizes end of line punctuation

##Usage

    b := normalize.Text([]byte(" this is .some ugLy TEXT. it’s not formatted very well at all    "))
    fmt.Println(string(b))
    // This is some ugly text. It's not formatted very well at all.

##See Also
  
  http://github.com/AlasdairF/Titlecase
