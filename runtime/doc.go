/*
Package runtime provides functionality for discovering and managing the PHP execution
environment required to interact with a Laravel project at runtime. It encapsulates
the process of locating a valid PHP binary and executing PHP code to retrieve essential
project information

By abstracting the PHP process management details, the runtime package allows client
code to seamlessly obtain runtime information without manually handling PHP code,execution or file discovery.
This design helps in maintaining a clear separation of concerns.

Key features:
  - Locate and validate the PHP executable for the target Laravel project.
  - Execute PHP code to extract project data.
  - Provide structured access to configuration and application binding data via repository types.

The package is intended to be used as a foundation for tools that need to analyze or interact
with Laravel projects by leveraging runtime information.
*/
package runtime
