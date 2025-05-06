# texml-go

This is a Go client library for [Telnyx Programmable Voice API](https://developers.telnyx.com/api/).

At the time of writing, Telnyx does not provide an official Go SDK for their APIs. Please check the official [Developer setup](https://developers.telnyx.com/docs/development#developer-setup) section of the Telnyx Docs for updates and a possible release of an official Go SDK in the future.

For now, this library only contains a package for generating [TeXML](https://developers.telnyx.com/docs/voice/programmable-voice/texml-fundamentals).

## ⚠️ WARNING! ⚠️ 

This library is still in active development. Be aware that not all TeXML verbs and nouns have been fully implemented yet. Please see below for a full list of the completed verbs and nouns. If you would like to contribute, please see feel free to submit a pull request.

## Installation

The recommended way to install `telnyx-go` is by using [Go modules](https://go.dev/ref/mod#go-get).

If you already have an initialized project, you can run the command below from your terminal in the project directory to install the library:

```shell
go get github.com/andersryanc/telnyx-go
```

If you are starting from scratch in a new directory, you will first need to create a go.mod file for tracking dependencies such as telnyx-go. This is similar to using package.json in a Node.js project or requirements.txt in a Python project. [You can read more about mod files in the Go documentation](https://golang.org/doc/modules/managing-dependencies). To create the file, run the following command in your terminal:

```shell
go mod init texml-example
```

Once the module is initialized, you may run the installation command from above, which will update your go.mod file to include telnyx-go.

## Completed Verbs and Nouns

- [x] [`<Dial>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial)
    - [x] [`<Number>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial#number-attributes)
    - [x] [`<Sip>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial#sip-attributes)
    - [x] [`<Queue>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/dial#queue-attributes)
- [x] [`<Conference>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/conference)
- [x] [`<Enqueue>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/enqueue)
- [x] [`<Gather>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/gather)
- [ ] [`<AIGather>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/aigather)
    - [ ] [`<Greeting>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/aigather#child-verbsnouns)
    - [ ] [`<Voice>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/aigather#child-verbsnouns)
    - [ ] [`<Parameters>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/aigather#child-verbsnouns)
    - [ ] [`<MessageHistory>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/aigather#child-verbsnouns)
- [x] [`<Hangup>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/hangup)
- [ ] [`<HttpRequest>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/httprequest)
    - [ ] [`<Request>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/httprequest#child-verbsnouns)
    - [ ] [`<Response>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/httprequest#child-verbsnouns)
- [x] [`<Leave>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/leave)
- [x] [`<Pause>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/pause)
- [x] [`<Play>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/play)
- [x] [`<Record>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/record)
- [x] [`<Redirect>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/redirect)
- [x] [`<Refer>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/refer)
    - [x] [`<ReferSip>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/refer#examples)
- [x] [`<Reject>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/reject)
- [x] [`<Say>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/say)
- [ ] [`<Siprec>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/siprec)
- [x] [`<Stop>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/stop)
- [x] [`<Stream>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/stream)
    - [x] [`<Start>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/stream)
- [x] [`<Suppression>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/suppression)
- [ ] [`<Transcript>`](https://developers.telnyx.com/docs/voice/programmable-voice/texml-verbs/transcription)
