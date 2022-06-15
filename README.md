# Autocomplete Web Service using GoLang

This is a simple test of an Autocomplete web service for Emapta-Zesty.io. The purpose is to return a list of Shakespeare's works based on the fragment that the user will provide or input. In this case, we will be using a query to simulate user's input. Shakespeare's works' list will be extracted from a JSON file. Each item on the JSON file contains only 2 itemms.

```JSON
{
    "Title": "Title of the work",
    "ReadCount": 0
}
```

The result is based on how frequent the user has accessed an item which is the `ReadCount`. If more than 1 item have equal frequency, the service will then lookup for the closest `Title` from the fragment provided. If still more than 1 item are equal, result will now be based alphabetically. Moreover, result is limited to 25 results only. This limit can be adjusted by modifying the constant value of the `resultLimit` inside main.go file.

Feel free to edit the JSON file. Update each `ReadCount` values to your likings.

You can also update the threshold for determining the amount of levenshtein distance by modifying the multiplier value of `thresholdMultiplier` inside main.go.

## Installation:

### Using VSCode:

1. Open a new window of VSCode.
2. Drag the folder 'autocomplete-test' to the newly opened VSCode window.
3. Open a new terminal (Testminal->New Terminal) or press Ctrl+Shift-`
4. Type `go run main.go` and hit Enter. If no errors, server is already running.
5. Copy and paste this URL to any browser: http://localhost:9000/autocomplete?term=th

### Using Docker:

1. Make sure Docker is running.
2. Make sure that the DockerFile file is included in the files.
3. In the terminal, type and run `docker build -t <NameYourImage> .`.
4. Then type `docker run -p 9000:9000 -tid <NameYourImage>`.
5. Copy and paste this URL to any browser: http://localhost:9000/autocomplete?term=th

### Using the build file

1. There are 2 files included, main.exe and main, for Windows and Linux respectively.
2. Make sure those files are in the root directory where the main.go resides.
3. Simply run the main.exe file. A console window will open.
4. Copy and paste this URL to any browser: http://localhost:9000/autocomplete?term=th

- Note that, I wasn't able to test this on a linux machine. Only on a linux docker image.

## Other ways to access the web service

### Using cURL:

1. Make sure the server is running.
2. Type in `curl http://localhost:9000/autocomplete?term=th` in a terminal window and hit Enter.

## Sample results

These are the results based on the test run that I did on my end:

### Fragment: "th"
    Twelfth Night
    Two Gentlemen of Verona
    Taming of the Shrew
    Tempest
    Timon of Athens
    Titus Andronicus
    Troilus and Cressida

### Fragment: "fr"
    Troilus and Cressida

### Fragment: "pi"
    King John
    King Lear
    Midsummer Night's Dream
    Pericles
    Richard II
    Richard III
    Timon of Athens
    Titus Andronicus
    Winter's Tale

### Fragment: "sh"
    None

### Fragment: "wu"
    Julius Caesar
    Much Ado about Nothing
    Winter's Tale

### Fragment: "ar"
    All's Well That Ends Well
    Antony and Cleopatra
    As You Like It
    Troilus and Cressida

### Fragment: "il"
    All's Well That Ends Well

### Fragment: "ne"
    Henry IV, Part I
    Henry IV, Part II
    Henry V
    Henry VI, Part I
    Henry VI, Part II
    Henry VI, Part III
    Henry VIII
    Measure for Measure
    Merchant of Venice
    Merry Wives of Windsor
    Pericles
    Tempest

### Fragment: "se"
    Henry IV, Part I
    Henry IV, Part II
    Henry V
    Henry VI, Part I
    Henry VI, Part II
    Henry VI, Part III
    Henry VIII
    Measure for Measure
    Merchant of Venice
    Merry Wives of Windsor
    Pericles
    Tempest
### Fragment: "pl"

    All's Well That Ends Well
    Pericles