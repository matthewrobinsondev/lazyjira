# Work In Progress!
On the quest to never leave the terminal.

## Setup and initialisation
To setup and connect to your jira account, follow the following instructions:

1. Setup a new personal access token (On the Jira UI)
```BASH
 Click settings cog > Atlassian account settings > Security > Personal Access Tokens
```

2. Setup cli to be able to authenticate and connect to your jira instance
```BASH
laztjira config auth 
```

**Config options (examples):**
- **Instance URI:** https://myapp.atlassian.net
- **Email:** email@example.com
- **Access Token:** XXXXX

## Running Lazy Jira
```BASH
lazyjira
```