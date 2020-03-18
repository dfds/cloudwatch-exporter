[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![Build Status](https://dev.azure.com/dfds/YourAzureDevOpsProject/_apis/build/status/Name-Of-CI-Pipeline?branchName=master)](https://dev.azure.com/dfds/YourAzureDevOpsProject/_build/latest?definitionId=1378&branchName=master)

<!-- TABLE OF CONTENTS -->
## Table of Contents

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)



<!-- ABOUT THE PROJECT -->
## About The Project

The CloudWatch Log exporter is used to export CloudWatch logs as Metrics by running as a pod inside of Kubernetes


### Built With

Tool list: 

* [Go][Go]
* [AWS SDK for Go][AWS-SDK]

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* Install Go on your development machine.

### Installation
 
1. Clone the repository

```sh
git clone https://github.com/dfds/cloudwatch-exporter.git
```

2. Change directory to cloudwatch-exporter
```sh
cd cloudwatch-exporter
```

3. Install project dependencies

```sh
go get -v
```


<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_



<!-- ROADMAP -->
## Roadmap

See the [open issues](https://github.com/github_username/repo/issues) for a list of proposed features (and known issues).



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/dfds/cloudwatch-exporter?style=plastic
[contributors-url]: https://github.com/dfds/cloudwatch-exporter/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/dfds/cloudwatch-exporter?style=plastic
[issues-url]: https://github.com/dfds/cloudwatch-exporter/issues
[license-shield]: https://img.shields.io/github/license/dfds/cloudwatch-exporter?style=plastic
[license-url]: https://github.com/dfds/cloudwatch-exporter/blob/master/LICENSE

[Go]: https://golang.org/
[AWS-SDK]: https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-cloudwatch-with-go-sdk.html
