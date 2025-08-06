# TabGo
This project is a Go application developed to open a list of URLs in a file in Chrome or Firefox browsers in batches. The URLs are opened in groups of a specified number (batch), and the user's approval is obtained between each group.

## Features
- Batch URL opening with Chrome or Firefox
- Automatic certificate error skipping for Chrome (--ignore-certificate-errors)
- Profile selection support for Firefox
- Batch opening and user confirmation after each batch

## Installation
1. Go must be installed: https://golang.org/dl/
2. Run it ```git clone https://github.com/barisbaydur/TabGo.git```
3. Compile it ```go build -o TabGo .```

## Usage
First, write the URLs you want to open into a file (default: webapps.txt) with one URL per line.

### Parameters
- ```file```: File containing the URL list (default: webapps.txt)
- ```browser```: Browser to be used (Chrome or Firefox)
- ```profile```: Firefox profile (for Firefox only)
- ```batch```: Number of URLs to be opened each time (default: 50)

## Examples
```TabGo.exe -file=myurls.txt -browser=chrome -batch=10```</br>
```TabGo.exe -browser=firefox -profile=Insecure```

## Notes
- To bypass certificate errors in Chrome, ```--ignore-certificate-errors``` is automatically added.
- It is not possible to automatically bypass certificate errors in Firefox. If necessary, you can create a special profile and change its settings. (If you are wondering how to create a profile for Firefox, you can find out by clicking [here](#how-to-create-a-firefox-profile).)
- You must press Enter to continue after each batch.
- <b>Make sure that the browser you are using has been added to PATH so that it can be run.</b>

# How To Create a Firefox Profile?
1. Open a terminal.
2. Run it ```firefox.exe -P```
3. Click to "Create a Profile" button. <b>(Save the path it provides during creation.)</b>
4. Go to the saved path.
5. Create a file named is ```user.js```.
6. Add the following lines to the end of the file.
```
user_pref("security.enterprise_roots.enabled", true);
user_pref("security.ssl.errorReporting.enabled", false);
user_pref("browser.ssl_override_behavior", 2);
user_pref("browser.xul.error_pages.enabled", false);
```

Note: These settings reduce some certificate warnings, but do not automatically bypass all warnings. Firefox does not allow you to completely bypass most certificate errors automatically for security reasons.
