# Autocomplete Test

This is a simple test of an Autocomplete web service for Emapta-Zesty.io. The purpose is to return a list of Shakespeare's works based on the fragment that the user will provide or input. In this case, we will be using a quary to simulate user's input. Shakespeare's works' list will be extracted from a JSON file. Each item on the JSON file contains only 2 itemms.

`{ Title: "Title of the work"\nReadCount: "How frequent this was accessed" }`

The result is based on how frequent the user have access an item which is the `ReadCount`. If more than 1 item have equal frequency, the service will then lookup for the closes Title from the fragment provided. If still more than 1 item are equal, result will now be based alphabetically.

## Installation:

### Using VSCode:

1. Open a new window of VSCode.
2. Drag the folder 'autocomplete-test' to the newly opened VSCode window.
3. Open a new terminal (Testminal->New Terminal) or press Ctrl+Shift-`
4. Type go run main.go and hit Enter. If no errors, server is already running.
5. Copy and paste this URL to any browser: http://localhost:9000/autocomplete?term=th

## Using Docker: