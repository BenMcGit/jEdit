# jEdit

<div id="top"></div>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/benmcgit/jedit">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center"></h3>
  <p align="center">
    Process, filter, and modify JSON in seconds
    <br />
    <a href="https://github.com/benmcgit/jedit"><strong>Explore the docs »</strong></a>
    <br />
    <a href="https://github.com/benmcgit/jedit/issues">Report Bug</a>
    ·
    <a href="https://github.com/benmcgit/jedit/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

As engineers we often find ourselves parsing through large amounts of data to uncover bugs in a service and to learn more about how a service operates. This process can be tedious and time consuming. 

jEdit is a tool that can be used to quickly process, filter, and modify large JSON datasets so you can get back to focusing on feature work, opposed to wasting time looking through logs. 

This tool was written in golang and can be used via the commandline. Continue reading to determine how to get started!

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Golang](https://go.dev/doc/)
* [Cobra](https://github.com/spf13/cobra)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

Use the following instructions to learn how to clone the jEdit repository and to build the jEdit binary.

### Prerequisites

* [Install Golang](https://go.dev/doc/install)
* Validate golang version is 1.7.2 or higher
  ```sh
  go version
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/benmcgit/jedit.git
   ```
2. Build jEdit binary
   ```sh
   go build
   ```
3. Validate the binary is available
   ```sh
   ./jedit --help
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Now that jEdit is installed, lets try it out! 

There is test data that we can execute these examples with. This test data is located in the [testdata](https://github.com/BenMcGit/jEdit/tree/master/testdata) directory.

### Reducing
```sh
cat testdata/yesterday.json | ./jedit query <FLAGS>
```

* Lets say we only want to return the records where the key "team" contains the value "team-x". Try out this...
  ```sh
  cat testdata/yesterday.json | ./jedit query --filter "team == team-x"
  ```
* Maybe we want to include all records except ones containing "team" equal to "team-x". Then we can use this...
  ```sh
  cat testdata/yesterday.json | ./jedit query --filter "team != team-x"
  ```
* We can use more than one filter to reduce the filter set even further. In this example we return all values with "team" equal to "team-x" and "severity" with a value of 4 or higher...
  ```sh
  cat testdata/yesterday.json | ./jedit query --filter "team == team-x" --filter "severity >= 4"
  ```
* The following filter operators can be used to reduce the dataset. These operators can be used on strings along with digits.
  ```sh
  cat testdata/yesterday.json | ./jedit query --filter "team < team-x" --filter "team != team-a" --filter "severity > 0" --filter "severity <= 4"
  ```
* If the key does not match any key in the record set, 0 entries will be returned...
  ```sh
  cat testdata/yesterday.json | ./jedit query --filter "thisdoesnotexist == team-x"
  ```

### Sorting
```sh
cat testdata/yesterday.json | ./jedit sort <KEY> <FLAGS>
```

* Use the "sort" command to sort the data by providing a key as an argument. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit sort severity
  ```
* By default, the sort command will sort the dataset in descending order. Use the "asc" flag to sort in ascending order. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit sort severity --asc
  ```
* You may also sort the dataset on non-numerical based key-value pairs. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit sort team
  ```

### Adding a new key-value pair
```sh
cat testdata/yesterday.json | ./jedit addKey <NEW_KEY> <NEW_VALUE> <FLAGS>
```

* There may be a reason to add a new key to each item in your dataset. You can accomplish this by providing the new key and new value as arguments to the "addKey" command. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit addKey newKey newValue
  ```
* Its possible that you only want to add the key to specific items in the dataset. You can use filters to accomplish this. You may use as many filters as you'd like. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit addKey newKey newValue --filter "team == team-x" --filter "severity > 3"
  ```
* What happens if the new key already exists? We may want to replace that data with a new value. To replace existing values we can use the "replace" flag. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit addKey severity 10 --replace
  ```
* By default, if the new key already exists in the dataset, the old value for that key will be retained. In this example, the output will not be modified:
  ```sh
  cat testdata/yesterday.json | ./jedit addKey severity 10
  ```
### Removing a key-value pair
```sh
cat testdata/yesterday.json | ./jedit removeKey <KEY> <FLAGS>
```

* We may want to remove a key-value pair to reduce complexity in our dataset even further. To accomplish this, we can use the "remove" command by supplying a key. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit removeKey team
  ```
* We can use filters to selectively remove objects from our dataset. If no filter is supplied, the key will be removed from all entries. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit removeKey team --filter "team == team-a"
  ```

### Modifying a key
```sh
cat testdata/yesterday.json | ./jedit modifyKey <KEY> <NEW_KEY_NAME> <FLAGS>
```

* There may be a situation where we want to add a prefix, suffix, or modify the name of a key altogether. In this scenario we can use the "modifyKey" command by supplying a key and new key as arguments. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit modifyKey team my_super_awesome_team
  ```
* Like the addKey, removeKey, and query commands, we can use filters to determine which objects in the dataset to modify. Example:
  ```sh
  cat testdata/yesterday.json | ./jedit modifyKey team my_super_awesome_team --filter "team == team-x" --filter "ts > 1642415085"
  ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Feature 1
- [ ] Feature 2
- [ ] Feature 3
    - [ ] Nested Feature

See the [open issues](https://github.com/benmcgit/jedit/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Ben McAdams - [@benwmcadams](https://twitter.com/benwmcadams) - mcadams.benj@gmail.com

Project Link: [https://github.com/benmcgit/jedit](https://github.com/benmcgit/jedit)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* []()
* []()
* []()

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/benmcgit/jedit/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/benmcgit/jedit/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/benmcgit/jedit/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/benmcgit/jedit/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/benmcgit/jedit/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png