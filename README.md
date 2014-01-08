ghost
=====

A simple command line gister.

Installation
============

* clone this repo
* run:

``` $ go install # Easiest is to create a working Golang setup with a GOPATH environment variable ```

* obtain a GitHub "Personal API Access Token" and save it in a file called '.gist' in your $HOME directory.
(On Windows, the home directory is usually C:\Users\<you>) 

* Personal API Access tokens are created on the GitHub website under "Account Settings".
* Account Settings -> Applications

Usage
=====

* Add 'ghost' to your system PATH.

```
$ ghost <file> [files ...] # unix
C:\> ghost <file> [files ...] # Windows
```
