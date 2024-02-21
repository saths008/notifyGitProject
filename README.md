# notifyGitProjectd

This is a Linux daemon that checks every hour whether a github repository has been pushed to. A notification is sent every hour with the relevant information.

## Set up

1. Install the golang toolchain.
2. In `cmd/notifyGitProjectd`, run `go build` to produce an executable named `notifyGitProjectd`.
3. Add a `.env` in `cmd/notifyGitProjectd` with the key `GH_TOKEN` and a GitHub personal access token with the requirement of reading user repositories.

To make this a linux daemon:

4. Create a `/etc/systemd/system/notifyGitProjectd.service`, with the contents:

```
[Unit]
Description=Notify Git Project

[Service]
ExecStart=<PATH-TO-PROJECT>/notifyGitProjectd/cmd/notifyGitProjectd/notifyGitProjectd <OWNER-OF-REPO> <REPO-NAME>
WorkingDirectory=<PATH-TO-PROJECT>/notifyGitProjectd/cmd/notifyGitProjectd/
Restart=always
User=<USERNAME>

[Install]
WantedBy=multi-user.target
```

5.

```bash
sudo systemctl daemon-reload
sudo systemctl start notifyGitProjectd
sudo systemctl enable notifyGitProjectd
sudo systemctl status notifyGitProjectd

```
