<div align="center">
  <a href="https://codeberg.org/tomkoid/supershush">
    <img src="assets/logo.png" alt="Logo" width="168" height="168">
  </a>

  <h3 align="center">SuperShush</h3>

  <p align="center">
    Stop audio from playing when the audio output changes
    <br />
    <a href="https://codeberg.org/tomkoid/supershush/issues/new">Report Bug</a>
    ·
    <a href="https://codeberg.org/tomkoid/supershush/issues/new">Request Feature</a>
  </p>
</div>

## What is SuperShush?

SuperShush is a daemon that is designed to prevent audio from continuing to play when the output device is changed, such as when switching from headphones to speakers.

## Installation

To stop audio playback when the output device is changed, SuperShush **requires** the `playerctl` package if the option in the config is enabled.

### From Source

To compile SuperShush from source, you'll need to have Go installed on your system. You can download the Go installer from the official Go website.

Once you have Go installed, run the following command to install SuperShush:

```bash
go install codeberg.org/tomkoid/supershush@0.2.0
```

## Running

To run supershush, simply execute the `supershush` command:

```bash
supershush
```

This will start the SuperShush daemon, which will monitor audio output changes and stop audio playback when necessary.

## Configuring

To configure SuperShush, you can edit the configuration file located at `~/.config/supershush/config.toml`. Here is an example configuration file:

```toml
# Resume audio after you change output back
# WARNING: This feature is not stable yet. The current implementation has big issues.
# Example: audio is playing on audio 1 and then you switch to
#          audio 2, the audio gets paused. after you switch
#          back to audio 1, the audio gets unmuted.
resume = false

# Enable MPC pausing 
mpc = false

# Enable playerctl stopping 
playerctl = true
```

## License
 
SuperShush is licensed under the MIT License. See the LICENSE file for more information.

## Contributing
 
If you'd like to contribute to SuperShush, please fork the repository and submit a pull request with your changes. You can also report issues or suggest new features on the issue tracker.
