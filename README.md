# gotokens

Need your code to authorise your connection to an API, or connect to a database? Of course you do.

Don't want to hard-code your password or auth-tokens into your codebase? Of course you don't.

So how do you proceed?

I found a solution online which involved setting environment variables which the authentication code could then read--but this raised more questions: what happens when we reboot? What if we want to use the program on more than one computer? Rather than setting the environment variables manually, creating an batch file or shell script to set them for us would be necessary--and then we're back to having them in a file on a drive somewhere.

Obviously this file could not be checked in to our code repository or we'd be right back where we started. And if we want several computers to have access to the tokens, we need the file to be located somewhere that is visible to all computers, but *not* visible to the world at large.

My preliminary solution was to create a `JSON` file on a LAN drive. My slightly more cunning solution was to turn that into a `git` repository (obviously not hosted remotely) and clone it onto each machine which needs access to the tokens (accessing the LAN drive could take several seconds if it had gone to sleep!) This has the benefit that I can have fast access (or even remote access if I take my laptop out with me) while retaining all the benefits of version control--most notably synchronisation--should I need to add more tokens. At the same time, my tokens are only visible to people with access to my LAN, which is good enough for me.

The original code I wrote to read information out of my Token file was quite specific to the application I was developing. `go-tokens` provides a more general interface that can be used as required.

## JSON File Structure

Currently I have *one* `JSON` file, containing my API credentials for interfacing with `Twitter`. The file is `Tokens/API/twitter.json` and it looks a little like this:

```json
{
  "tokens": [
    {
      "name": "Twitter_App_Name",
      "credentials": [
        {
          "KEY1": "abcde12345",
          "SECRET1": "abcde12345",
          "KEY2": "abcde12345",
          "SECRET2": "abcde12345"
        }
      ]
    }
  ]
}
```

Other possible `credential` names might be "USERNAME" and "PASSWORD" or whatever your particular application required. If connecting to a database it might include "HOST" or "SCHEMA"--although that would probably then live in `Tokens/DB/mysql.json`.

## Using gotokens

To use this information in your code, you would need something like:

```go
import "github.com/pjsoftware/gotokens"

token = new(gotokens.Token)
token.Search({"..", "//lan/shared/gitfiles"})
token.Import("Tokens/API/twitter.json", "Twitter_App_Name")
key1 := token.Credential("KEY1")
sec1 := token.Credential("SECRET1")
```

This will, of course, all change as I develop the code.
