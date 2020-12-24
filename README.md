# devstats
A small application that gathers code review and build metrics.

This is a little light on features and test coverage at the moment, but it'll come around.  For now, it gathers code review open time for a 28-day period from GitHub and Crucible.

## Configuration:

Add a file like this to the application root.

Note:  this manner of configuration is an initial iteration.  It will be deprecated in favor of a more traditional CLI flow soon.

**config.yml**
```yaml
github:
  access:
    key: "your-github-api-key"
    baseurl: "https://<github url>/api/v3"
  repositories:
    - repository: "pex"
      branch: "master"
      owner: "pex"
    - repository: "help"
      branch: "master"
      owner: "UIC"
crucible:
  access:
    user: your.ad.username
    key: your.ad.password
    baseurl: "https://<your crucible url>/rest-service/reviews-v1/filter/details.json"
  projects:
    - name: "CR-EXT"
```
## Running the Application

1. Clone this repository.
1. Configure the application as above.
1. `go build`
1. `go install`
1. `devstats prstats` 
