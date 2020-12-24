package crucible

const ReviewLinkBase = "https://crucible.workday.com/cru/"
const CrucibleTimeFormat = "2006-01-02T15:04:05-0700"

type CrucibleComments struct {
	Comments []map[string]interface{}
}

type PermaId struct {
	Id string
}

type CrucibleReview struct {
	ProjectKey      string
	PermaId         PermaId
	CreateDate      string
	CloseDate       string
	JiraIssueKey    string
	GeneralComments CrucibleComments
}

type CrucibleProjectReviews struct {
	DetailedReviewData []CrucibleReview
}

/**
Sample Response:

{
  "detailedReviewData" : [ {
    "projectKey" : "CR-FOO",
    "name" : "Example review.",
    "description" : "Description or statement of objectives for this example review.",
    "author" : {
      "userName" : "auth",
      "displayName" : "Jane Authy",
      "avatarUrl" : "http://foo.com/avatar"
    },
    "moderator" : {
      "userName" : "scott",
      "displayName" : "Scott the Moderator",
      "avatarUrl" : "http://foo.com/avatar"
    },
    "creator" : {
      "userName" : "joe",
      "displayName" : "Joe Krustofski",
      "avatarUrl" : "http://foo.com/avatar"
    },
    "permaId" : {
      "id" : "CR-FOO-21"
    },
    "summary" : "some review summary.",
    "state" : "Review",
    "type" : "REVIEW",
    "allowReviewersToJoin" : true,
    "metricsVersion" : 4,
    "createDate" : "2020-09-11T10:53:52.352+0000",
    "dueDate" : "2020-09-14T10:53:52.352+0000",
    "reminderDate" : "2020-09-12T10:53:52.352+0000",
    "linkedIssues" : [ "DEF-456", "ABC-123", "GHI-789" ],
    "reviewItems" : {
      "reviewItem" : [ {
        "permId" : {
          "id" : "CFR-69"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir68/a.txt",
        "fromRevision" : "69",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5691/dir68/a.txt",
        "toPath" : "dir68/a.txt",
        "toRevision" : "691",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/6560/dir68/a.txt",
        "patchUrl" : "/contextcru/CR-FOO-21/downloadpatch/6629/patch-851.patch",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : false,
        "anchorData" : {
          "anchorPath" : "trunk",
          "anchorRepository" : "REPO",
          "stripCount" : 0
        },
        "commitDate" : 1599821632413,
        "expandedRevisions" : [ ]
      }, {
        "permId" : {
          "id" : "CFR-70"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir69/d.txt",
        "fromRevision" : "70",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5694/dir69/d.txt",
        "toPath" : "dir69/d.txt",
        "toRevision" : "701",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/6275/dir69/d.txt",
        "patchUrl" : "/contextcru/CR-FOO-21/downloadpatch/6537/patch-5413.patch",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : false,
        "anchorData" : {
          "anchorPath" : "branches/branch1",
          "anchorRepository" : "REPO",
          "stripCount" : 0
        },
        "commitDate" : 1599821632415,
        "expandedRevisions" : [ ]
      }, {
        "permId" : {
          "id" : "CFR-71"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir70/c.txt",
        "fromRevision" : "71",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5697/dir70/c.txt",
        "toPath" : "dir70/c.txt",
        "toRevision" : "711",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/6043/dir70/c.txt",
        "patchUrl" : "/contextcru/CR-FOO-21/downloadpatch/5917/patch-5963.patch",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : false,
        "anchorData" : {
          "anchorPath" : "branches/branch1",
          "anchorRepository" : "REPO",
          "stripCount" : 0
        },
        "commitDate" : 1599821632416,
        "expandedRevisions" : [ ]
      }, {
        "permId" : {
          "id" : "CFR-72"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir71/b.txt",
        "fromRevision" : "72",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5700/dir71/b.txt",
        "toPath" : "dir71/b.txt",
        "toRevision" : "721",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5800/dir71/b.txt",
        "patchUrl" : "/contextcru/CR-FOO-21/downloadpatch/6274/patch-7579.patch",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : true,
        "anchorData" : {
          "anchorPath" : "tags/tag1",
          "anchorRepository" : "REPO",
          "stripCount" : 0
        },
        "commitDate" : 1599821632417,
        "expandedRevisions" : [ ]
      }, {
        "permId" : {
          "id" : "CFR-73"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir72/d.txt",
        "fromRevision" : "73",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5702/dir72/d.txt",
        "toPath" : "dir72/d.txt",
        "toRevision" : "731",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/6531/dir72/d.txt",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : true,
        "commitDate" : 1599821632417,
        "expandedRevisions" : [ ]
      }, {
        "permId" : {
          "id" : "CFR-74"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir73/e.txt",
        "fromRevision" : "74",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5704/dir73/e.txt",
        "toPath" : "dir73/e.txt",
        "toRevision" : "741",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/6616/dir73/e.txt",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : true,
        "commitDate" : 1599821632417,
        "expandedRevisions" : [ ]
      }, {
        "permId" : {
          "id" : "CFR-75"
        },
        "participants" : [ {
          "user" : {
            "userName" : "fred",
            "displayName" : "Fred Eaglesmith",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : true
        }, {
          "user" : {
            "userName" : "clund",
            "displayName" : "Corb Lund",
            "avatarUrl" : "http://foo.com/avatar"
          },
          "completed" : false
        } ],
        "repositoryName" : "REPO",
        "fromPath" : "dir74/f.txt",
        "fromRevision" : "75",
        "fromContentUrl" : "/context/cru/CR-FOO-21/rawcontent/5706/dir74/f.txt",
        "toPath" : "dir74/f.txt",
        "toRevision" : "751",
        "toContentUrl" : "/context/cru/CR-FOO-21/rawcontent/6434/dir74/f.txt",
        "fileType" : "File",
        "commitType" : "Modified",
        "authorName" : "evzijst",
        "showAsDiff" : true,
        "commitDate" : 1599821632417,
        "expandedRevisions" : [ ]
      } ]
    },
    "generalComments" : {
      "comments" : [ {
        "message" : "I thought you said you were going to *remove* this line?",
        "draft" : false,
        "deleted" : false,
        "defectRaised" : false,
        "defectApproved" : false,
        "readStatus" : "UNREAD",
        "user" : {
          "userName" : "joe",
          "displayName" : "Joe Bloggs",
          "avatarUrl" : "http://foo.com/avatar"
        },
        "createDate" : "2020-09-11T10:53:52.392+0000",
        "permaId" : {
          "id" : "CR:1"
        },
        "messageAsHtml" : "I thought you said you were going to <b>remove</b> this line?",
        "parentCommentId" : { }
      } ]
    },
    "transitions" : {
      "transitionData" : [ {
        "name" : "action:summarizeReview",
        "displayName" : "Summarize"
      } ]
    },
    "actions" : {
      "actionData" : [ {
        "name" : "action:abandonReview",
        "displayName" : "Abandon"
      }, {
        "name" : "action:approveReview",
        "displayName" : "Approve"
      }, {
        "name" : "action:commentOnReview",
        "displayName" : "Comment"
      }, {
        "name" : "action:completeReview",
        "displayName" : "Complete"
      }, {
        "name" : "action:modifyReviewFiles",
        "displayName" : "Edit Review Details"
      } ]
    },
    "jiraIssueKey" : "ABC-123"
  } ]
}
*/
