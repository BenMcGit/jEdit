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
    <a href="https://pkg.go.dev/github.com/benmcgit/jedit"><strong>Explore the docs »</strong></a>
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
* Validate golang version is 1.17.2 or higher
  ```sh
  go version
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/benmcgit/jedit.git
   ```
2. Build jEdit binary (in root directory)
   ```sh
   cd jedit && ./scripts/build.sh
   ```
3. Validate the binary is available
   ```sh
   ./jedit --help
   ```
4. Execute the test script. This script covers the following scenario:
   * We have a list of json records (`./testdata/yesterday.json`) and want to add a new key-vale pair to them.
   * We only want to add the key-value pair if the json record contains the key "team" and this key contains the value "team-x". 
   * The new key is named "incident_id" and the value for that key will be "6502".
   ```sh
   ./scripts/test.sh
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Now that jEdit is installed, lets try it out! 

There are two ways jEdit can be consumed
1. Command Line
2. Public API

See below examples of how to use each below.

### Command Line

There is test data that we can execute these examples with. This test data is located in the [testdata](https://github.com/BenMcGit/jEdit/tree/master/testdata) directory.

#### Reducing
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

#### Sorting
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

#### Adding a new key-value pair
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
#### Removing a key-value pair
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

#### Modifying a key
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

#### Providing input

There are two ways a user can provide input. 
1. Piping via Stdin
  ```sh
  cat testdata/yesterday.json | ./jedit modifyKey ...
  ```
2. Using the `--input` flag
  ```sh
  ./jedit modifyKey ... --input testdata/yesterday.json
  ```

#### Defining output location

By default, jEdit will print its output to Stdout. Optianlly you can use the `--output` to create a new file and write to it. 
  ```sh
  ./jedit modifyKey ... --output my_output_file.json
  ```

#### Need more help?
```sh
cat testdata/yesterday.json | ./jedit --help

Parsing and editing JSON in bulk can be time-consuming
and difficult. jEdit aims to help engineers save time by providing them 
a tool that can reduce, filter, and modify their existing JSON dataset.

Usage:
  jedit [command]

Available Commands:
  addKey      Adds an additional key to object(s) in your dataset
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  modifyKey   Modifies existing keys on object(s) in your dataset
  query       Reduces the amount of object(s) in your dataset
  removeKey   Removes existing keys on object(s) in your dataset
  sort        Sorts objects in your dataset base on a user-provided key

Flags:
  -h, --help            help for jedit
      --input string    Path to JSON file to be parsed
      --output string   Path to file to write resulting JSON to. If not existent, it will be created.

Use "jedit [command] --help" for more information about a command.
```

### Public API

Please see the documentation for how to use jedit within your golang project [here](https://pkg.go.dev/github.com/benmcgit/jedit). 

IMPORTANT: jEdit must be consumed as a go module (required after go 1.17).You can create a go-module by running `go mod init <YOUR_MODULE_NAME>` (Please be sure to run `go get -v` to update your dependencies).

Example:
```go
package main

import (
	"log"

	"github.com/benmcgit/jedit/pkg/jedit"
)

func main() {
	filtersStr := []string{"severity >= 4"}
	filePath := "yesterday_reduced.json"

	// validates input and creates an instance of Logs
	logs, err := jedit.ParseFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// validates input and creates an array of Filter instances
	filters, err := jedit.ParseFilters(filtersStr)
	if err != nil {
		log.Fatal(err)
	}

	// apply commands on the logs
	logs.Add("urgency", "HIGH", false, filters)

  // write to stdout
	logs.Print()

  // write to a file
  logs.WriteToFile("example_output.json")
}
```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

### Nice-to-haves
- [ ] Preformance Improvements
    - [ ] Use concurrent go-threads to parse stdin
- [ ] Allow user to OR filters together (opposed to ANDing them together)
- [ ] Add support for multi-line JSON
- [ ] Improved validation around supplying a non-JSON based file

### Required for first release
- [x] Consume input stream from
    - [x] a terminal stdin
    - [x] a specified input source
- [x] Output data to
    - [x] a terminal stdout
    - [x] a specified output destination
- [x] Provide a way to compare values (less or greater than, longer or shorter than, …) rather than just equality or difference. 
- [x] Provide a way to apply operations only on objects only if they match a given predicate.
- [x] Rejecting an object based on the value of a specific field. 
- [x] Retaining an object only if it has a field with a specific value. 
- [x] Adding a key value pair on an object
- [x] Prefixing a key with a string
- [x] Provide a proper README.md

See the [open issues](https://github.com/benmcgit/jedit/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- TAKEAWAYS -->
## Project Takeaways

Besides creating a tool that worked functionally, I put a lot of thought into two non-functional requirements - Ease of use, preformace.

To make sure jEdit was easy to use, I went through a few design iterations to determine a way I could use a similar filtering process across different commands. This led me to the design where I use a the same filter flag process across the entire application (Ex. `--filter "key == value"`). I hoped that this would allow for my users to learn the format once, and be able to apply the same logic across the query, addKey, removeKey, and modifyKey commands. The tradoff here is a good amount of validation for each filter. 

One consideration I was debating on is whether I should AND all the filters together, or OR all of the filters together. I decided AND was more natural from a users perspective so I went with that. I think a cool new feature would be to allow for OR using a flag so the user can decide. I added this to the future improvements section as a nice-to-have. 

As far as preformance, I believe this can be improved. My current logic is to first read in the JSON file and store each record as a type of "Log". After this, I continue to fullfill the customers requested command based on the filters provided. There is an opportunity to filter or modify each record as its being read into momory. If preformance became an issue, this is where I would start investigating for potential improvements. 

This project was my first experiance using a tool like [Cobra](https://github.com/spf13/cobra). This tool allowed me to build commands that contain arguments and flags with simplicity. Along with being easy, it saved me a lot of time when it came to generating the jEdit user manual (`./jedit --help`). The next time you build a commandline tool using Golang, I would highly recommend checking out Cobra. 

While I have a lot of experiance with CI tooling, I have not used GitHuib Actions very much. It was really easy to setup for a small project like this. Take a look at the [Actions Tab](https://github.com/BenMcGit/jEdit/actions) to learn more. Right now, my CI pipeline is setup to run after every commit to master. 

For fun, I used a free online service that suggested and generated a logo for jEdit. I was pleasently suprised how easy it was to use so I added those images to the [images](https://github.com/BenMcGit/jEdit/tree/master/images) directory. If given more time I would try to come up with a logo myself, but I appreciate [Hatchful](https://hatchful.shopify.com/) for making my life a little easier and jEdit a little more official :). 

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

* [Hatchful by Shopify](https://hatchful.shopify.com/)
* [My beautiful fiancée for her endless support :)](https://www.zola.com/wedding/benandchrista2022)

<p align="right">(<a href="#top">back to top</a>)</p>
