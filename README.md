AudStopper
What is AudStopper?
AudStopper is a daemon that stops audio from playing when the audio output changes. It's designed to prevent audio from continuing to play when the output device is changed, such as when switching from headphones to speakers.

Installation Guide
Compiling from Source
To compile AudStopper from source, you'll need to have Go installed on your system. You can download the Go installer from the official Go website.

Once you have Go installed, clone the AudStopper repository:

git clone https://codeberg.org/tomkoid/audstopper.git
Change into the cloned repository and run the following command to build AudStopper:

go build main.go
This will create an executable file called audstopper in the current directory.

Installing the Binary
If you don't want to compile AudStopper from source, you can download the pre-compiled binary from the releases page.

Once you've downloaded the binary, you can install it to a directory of your choice. For example, to install it to /usr/local/bin, run the following command:

sudo install audstopper /usr/local/bin
Running AudStopper
To run AudStopper, simply execute the audstopper command:

audstopper
This will start the AudStopper daemon, which will monitor audio output changes and stop audio playback when necessary.

License
AudStopper is licensed under the MIT License. See the LICENSE file for more information.

Acknowledgments
AudStopper was inspired by the Audacious audio player and its plugins. The project's code is influenced by the RPM spec file for Audacious, which can be found in the Fedora repository.

Contributing
If you'd like to contribute to AudStopper, please fork the repository and submit a pull request with your changes. You can also report issues or suggest new features on the issue tracker.

go install codeberg.org/tomkoid/audstopper@latest
