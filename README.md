# AudStopper

## What is AudStopper?

AudStopper is a daemon that is designed to prevent audio from continuing to play when the output device is changed, such as when switching from headphones to speakers.

# Installation
 
## From Source

To compile AudStopper from source, you'll need to have Go installed on your system. You can download the Go installer from the official Go website.

Once you have Go installed, run the following command to install audstopper:

```bash
go install codeberg.org/tomkoid/audstopper@0.1.0
```

# Running

To run AudStopper, simply execute the audstopper command:

```bash
audstopper
```

This will start the AudStopper daemon, which will monitor audio output changes and stop audio playback when necessary.

# Configuring

To configure AudStopper, you can edit the configuration file located at `~/.config/audstopper/config.toml`. Here is an example configuration file:

```toml
# Enable MPC pausing 
mpc = false

# Enable playerctl stopping 
playerctl = true
```

# License
 
AudStopper is licensed under the MIT License. See the LICENSE file for more information.

# Contributing
 
If you'd like to contribute to AudStopper, please fork the repository and submit a pull request with your changes. You can also report issues or suggest new features on the issue tracker.
