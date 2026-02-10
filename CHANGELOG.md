# Changelog

All notable changes to this project will be documented in this file. See [conventional commits](https://www.conventionalcommits.org/) for commit guidelines.

---
## [0.0.13](https://github.com/laravel-ls/laravel-ls/src/v0.0.13) - 2026-02-10

### Miscellaneous

- use github.com/laravel-ls/protocol instead of internal package ([bf94ecd](https://github.com/laravel-ls/laravel-ls/commit/bf94ecd8a76727df43a30b9a360097e2c3cf45d6))
- remove -p flag as the compiler will use all cores per default. ([3bfbcd0](https://github.com/laravel-ls/laravel-ls/commit/3bfbcd06298efa900256530e27b3025e7d8644e8))
- show completion on empty string ([6b4ffa1](https://github.com/laravel-ls/laravel-ls/commit/6b4ffa17bfe8535f35a60c79042ab8aedb417bcc))
- variable renaming to keep it consistant with other providers ([ce51aae](https://github.com/laravel-ls/laravel-ls/commit/ce51aaeb030caf4ff68f15ef5f080882ae2fc0b2))
- adding code action to create non-existant route ([7dfd748](https://github.com/laravel-ls/laravel-ls/commit/7dfd7483c6511b49945534175a5913f21c0ef301))
- update implemented features ([e055004](https://github.com/laravel-ls/laravel-ls/commit/e055004f9b7016002c1837a46bc1681f83b79b0f))
- some cleanup and refactoring ([ff31e10](https://github.com/laravel-ls/laravel-ls/commit/ff31e10092460b4a4fde95ce88879ab3ec7c6b20))
- typo fix ([15015d7](https://github.com/laravel-ls/laravel-ls/commit/15015d727733c7cae0b4b7a37c38528097c3b6e2))
- project should have a reference to the runtime.Process interface, not runtime.PHPProcess directly. ([55b1355](https://github.com/laravel-ls/laravel-ls/commit/55b1355a8122b26125c7f5b7f94ea80ea77e4c28))
- document php requirement ([18f434b](https://github.com/laravel-ls/laravel-ls/commit/18f434bb4c0198fea839928117e875ec516df552))
- document mason ([21efb52](https://github.com/laravel-ls/laravel-ls/commit/21efb526eae8078efae191e302725b4442e42ee2))
- set diagnostics to Warning severity instead of Error. less dramatic :) ([bd35c38](https://github.com/laravel-ls/laravel-ls/commit/bd35c3889db874884e6e0dee58029a7fcd758d02))
- update laravel-ls/protocol ([ecdef55](https://github.com/laravel-ls/laravel-ls/commit/ecdef5518ad25c0e98b74cafac890cd512edada9))
- update pnx/tree-sitter-dotenv ([aa507fc](https://github.com/laravel-ls/laravel-ls/commit/aa507fc2557b7d0a74a33f75690988dd565c57ae))
- Revert "go.mod: update pnx/tree-sitter-dotenv" ([31f7679](https://github.com/laravel-ls/laravel-ls/commit/31f7679c3199e941ee0cab3d65c56e38169b954a))
- MacOS is officially supported now. remove the old text. ([3adb1ce](https://github.com/laravel-ls/laravel-ls/commit/3adb1ce6df894edb5e58aa17e4c2b96ae38761f4))

---
## [0.0.12](https://github.com/laravel-ls/laravel-ls/src/v0.0.12) - 2025-10-21

### Miscellaneous

- run tests on MacOS ARM64 ([26e96e6](https://github.com/laravel-ls/laravel-ls/commit/26e96e6bd7666775f92f4288d099c22149fe8a54))
- fix to correctly run on macos arm. ([c90d591](https://github.com/laravel-ls/laravel-ls/commit/c90d5912e5879c4dfb3958bd21891a1a436ac487))
- fix test again ([24c3bb8](https://github.com/laravel-ls/laravel-ls/commit/24c3bb86a709cbf89f53efe01ea16146b0acfb8d))
- remove macos-15-intel@arm64 ([7d73e7a](https://github.com/laravel-ls/laravel-ls/commit/7d73e7a3ca46c50f13d4feb4e85e81d32b39f5fa))
- update versions ([29839c2](https://github.com/laravel-ls/laravel-ls/commit/29839c20451551970286a116907eadb043996e5b))
- in TypeByFilename() have all checks use the lookup table by having test functions, filetype pairs. ([6f867e9](https://github.com/laravel-ls/laravel-ls/commit/6f867e9d750bd468b1a17a3727b01f50e5eb3c77))
- improve env file detection, should only match ".env" or filenames starting with ".env." ([17ed66e](https://github.com/laravel-ls/laravel-ls/commit/17ed66ef62ea0175919d0ef52eae1d58ceadb484))
- add goimports and add extra rules for gofumpt ([7212ba6](https://github.com/laravel-ls/laravel-ls/commit/7212ba6a227f65b8772a9c1cf9119d0af0a6a2d4))
- linting fixes. ([a050780](https://github.com/laravel-ls/laravel-ls/commit/a050780be9053c304ebbf3ae01dd958e83d26c7e))
- build for macos arm64 ([0213088](https://github.com/laravel-ls/laravel-ls/commit/02130881fb2ecd5193dc3566f27053748e63b14b))

---
## [0.0.11](https://github.com/laravel-ls/laravel-ls/src/v0.0.11) - 2025-10-20

### Miscellaneous

- add URL:asset() ([efeff67](https://github.com/laravel-ls/laravel-ls/commit/efeff67e92e2e14a2ec187aa054e6cd88d53f744))
- update readme regarding windows and MacOS support ([a93641a](https://github.com/laravel-ls/laravel-ls/commit/a93641a91cda5ecba07e402a8824996a7f230eb5))
- add install.sh ([19ab212](https://github.com/laravel-ls/laravel-ls/commit/19ab2124e4f94f2d4853e019fe67335833706955))
- try running tests on macos ([07afd03](https://github.com/laravel-ls/laravel-ls/commit/07afd0331ce5c8f7e98a6bf892b6ea200cef40aa))
- try build release for macos ([f103074](https://github.com/laravel-ls/laravel-ls/commit/f1030740bb5c1418f24ca80d6722baa7b625fcbe))
- github workflows: update setup-go to v6 ([0bfa8ed](https://github.com/laravel-ls/laravel-ls/commit/0bfa8ed530781f3f6f808fe13806742d50876ae8))
- update required go version to 1.23 ([94aca8b](https://github.com/laravel-ls/laravel-ls/commit/94aca8ba2c82f77a9ac9614adf9005eb36e21aed))

---
## [0.0.10](https://github.com/laravel-ls/laravel-ls/src/v0.0.10) - 2025-08-01

### Miscellaneous

- Support macOs for the make file ([03912b1](https://github.com/laravel-ls/laravel-ls/commit/03912b12980ac87063f5eaa4af677ba24494463b))
- Support macOs for the make file ([#12](https://github.com/laravel-ls/laravel-ls/issues/12)) ([fbf3fee](https://github.com/laravel-ls/laravel-ls/commit/fbf3fee7b404d26f72916a49459281e81a14ea68))
- in HandleInitialize() set DiagnosticProvider field correctly. ([0040f75](https://github.com/laravel-ls/laravel-ls/commit/0040f7525a3f7007c78d239cfb1b3530b6c8bad0))
- Merge branch 'lsp-initialization-diagnostics-provider' ([eb43683](https://github.com/laravel-ls/laravel-ls/commit/eb436836c478a4e676d63d6424a12bbca48b4774))

---
## [0.0.9](https://github.com/laravel-ls/laravel-ls/src/v0.0.9) - 2025-07-01

### Bug Fixes

- fix windows path in tests ([4a9ef42](https://github.com/laravel-ls/laravel-ls/commit/4a9ef42e5d11eefb0935727e049264361efdd0e1))

### Miscellaneous

- PHPProccess should be PHPProcess ([63e676f](https://github.com/laravel-ls/laravel-ls/commit/63e676f704e523f5b9fff8c2f629e118f0445276))
- add Process interface ([9c30f49](https://github.com/laravel-ls/laravel-ls/commit/9c30f49b78537808778551b8d92ec10a8169efb4))
- make CallScript take a object that implements the Process interface. ([37c8167](https://github.com/laravel-ls/laravel-ls/commit/37c816728d581e6314c75a1e1966005c621ac86c))
- in CallScript() rename call variable to proc. ([c049f87](https://github.com/laravel-ls/laravel-ls/commit/c049f8741c42d0b34deba3485273567d61cb64c6))
- Adding runtime/script_test.go ([62212cb](https://github.com/laravel-ls/laravel-ls/commit/62212cb62af15b3248a4d732cfc529de50c07498))
- in CallScript() make it abit more clear if there is a json decoding error. ([8239b45](https://github.com/laravel-ls/laravel-ls/commit/8239b45aea51d90cfa5299ed284241b0ca65ac77))
- run tests on windows. ([fda635c](https://github.com/laravel-ls/laravel-ls/commit/fda635c1c942cdebe7aa891d2dd770e68a6914ee))
- build release for windows. ([ca3664e](https://github.com/laravel-ls/laravel-ls/commit/ca3664e35300a6151a0539d14449ce1b24b8a598))
- wrap $GITHUB_ENV in quotes. ([2a8e463](https://github.com/laravel-ls/laravel-ls/commit/2a8e463be5b5a1a2d850e9f8bdf70031eb08c5ea))
- split tests into windows and non windows files ([6062d65](https://github.com/laravel-ls/laravel-ls/commit/6062d65205e9742750b392be5c224dbeb030c920))
- split tests into windows and non windows files ([4cc94e8](https://github.com/laravel-ls/laravel-ls/commit/4cc94e89d39d6868bd8677495e614ec0fb470ee3))
- Initial setup for routes support ([7c60da1](https://github.com/laravel-ls/laravel-ls/commit/7c60da1a22b589edd3a45be885089f7f6d193871))
- create temporary scripts in vendor/_laravel-ls directory ([ac34bb2](https://github.com/laravel-ls/laravel-ls/commit/ac34bb2c666bb53d07e3e6f17b15267fbb1304df))
- Update runtime/process.go ([a2f0e9b](https://github.com/laravel-ls/laravel-ls/commit/a2f0e9bb0d36cc83d0ad11af54347989a41b7192))
- Support cancelRequest, shutdown, and exit ([26727b6](https://github.com/laravel-ls/laravel-ls/commit/26727b631bca0659eea60b1d4369571486337f98))
- Support cancelRequest, shutdown, and exit ([#9](https://github.com/laravel-ls/laravel-ls/issues/9)) ([175b1d4](https://github.com/laravel-ls/laravel-ls/commit/175b1d4dc585f9044ad18f22598ee6e9bc92eeb4))
- make sure LogFile is open during the last call to the logger ([6b1c5dc](https://github.com/laravel-ls/laravel-ls/commit/6b1c5dcb0ee9a9b9dcca9b55443d2557f8b9b282))
- close connection instead of calling os.Exit() ([b115655](https://github.com/laravel-ls/laravel-ls/commit/b11565584209b266190a63f62e26c32893f16f38))
- return exit code from Run() ([9d056f7](https://github.com/laravel-ls/laravel-ls/commit/9d056f7112edafc2cdd1b46348c7199a86799cd9))
- call os.Exit() with code from cmd.Run() ([f416df2](https://github.com/laravel-ls/laravel-ls/commit/f416df240e4d2a800c20f444070c068c10b4322a))
- implement exit notification correctly ([c69f04f](https://github.com/laravel-ls/laravel-ls/commit/c69f04f69d7fcc2da7fcfddc06eacb5eb2506f17))
- minor fix. ([e9d0d19](https://github.com/laravel-ls/laravel-ls/commit/e9d0d194c673efab9b52c41efed4ce5c73c86b30))
- sail needs to be executed from the root directory to work correctly. ([fa390d8](https://github.com/laravel-ls/laravel-ls/commit/fa390d829f9bc7274913346c708c6c3a98cd17ce))
- update to version 2.1.6 ([5fe68bb](https://github.com/laravel-ls/laravel-ls/commit/5fe68bb1aa7b21133a6902bce913effe7ac20bd2))
- ignore test files ([118fa2e](https://github.com/laravel-ls/laravel-ls/commit/118fa2e93c69587c74b292c380207467cc3a7ed0))
- explicit ignore error returned from afero.Walk() to tell the linter that this is not a mistake ([83dcd8f](https://github.com/laravel-ls/laravel-ls/commit/83dcd8f3701f8c4369d2462ae3a7a54b846697af))
- explicit ignore error returned from afero.Walk() to tell the linter that this is not a mistake ([60426f3](https://github.com/laravel-ls/laravel-ls/commit/60426f385308631d8c1d627020b6c81d8f21b945))
- adding more badges and update the test badge to use shields.io ([aefc880](https://github.com/laravel-ls/laravel-ls/commit/aefc880e7c2475964db5cf03ca9c94c7ae6d88a0))
- make go version badge use the same style as the rest. ([f0a34d3](https://github.com/laravel-ls/laravel-ls/commit/f0a34d3d86978d1096932236b64cc4b93f1e1d74))
- use shield.io's link parameter for badges. ([7d9eef3](https://github.com/laravel-ls/laravel-ls/commit/7d9eef3567a51d5ce09534fda82630b90ac8339e))
- Revert "README.md: use shield.io's link parameter for badges." ([c6a2ffd](https://github.com/laravel-ls/laravel-ls/commit/c6a2ffdc7a640f7110c42cca053c3a67895f1748))
- fix badges. ([e515720](https://github.com/laravel-ls/laravel-ls/commit/e515720698ab17ee9e3795648ccf61640fde5218))
- Adding CONTRIBUTING.md ([96daa5a](https://github.com/laravel-ls/laravel-ls/commit/96daa5a8fdc43c41a08cfb663968169a1a896230))
- Merge branch 'main' into support-routes ([216200f](https://github.com/laravel-ls/laravel-ls/commit/216200fcb9dea98f5d6fff3eb981899275a3201e))
- Formatting ([a8999c4](https://github.com/laravel-ls/laravel-ls/commit/a8999c4c20a64e65f22dda7324de4f51b3738272))
- Update laravel/providers/route/provider.go ([8507e05](https://github.com/laravel-ls/laravel-ls/commit/8507e055d3dcd5de60ce5175c2ef54f97bf82549))
- Formatting ([2486442](https://github.com/laravel-ls/laravel-ls/commit/2486442bb84b083d05ce2005fbc3860ce703e09f))
- Initial setup for routes support ([#10](https://github.com/laravel-ls/laravel-ls/issues/10)) ([934a19c](https://github.com/laravel-ls/laravel-ls/commit/934a19c41bf260c106cdfba9daf93803e465cf5a))

---
## [0.0.7](https://github.com/laravel-ls/laravel-ls/src/v0.0.7) - 2025-06-24

### Miscellaneous

- fix log message ([365dfd8](https://github.com/laravel-ls/laravel-ls/commit/365dfd8be64b68610d4ab9e23c206807a4420469))
- update with nvim-lspconfig documentation ([0e00788](https://github.com/laravel-ls/laravel-ls/commit/0e00788a993939bc1ee524ea0f508feaf541d188))
- update with more complete documentation and types from spec. ([06bab27](https://github.com/laravel-ls/laravel-ls/commit/06bab27d5e508f8d956929e1cfcafc073b9ea5c0))
- fix step name ([77c3356](https://github.com/laravel-ls/laravel-ls/commit/77c33569437b29dad1ba5efe04f5b7a4f371dd74))
- check viper.ConfigFileUser() and report no file used if empty. ([f92475d](https://github.com/laravel-ls/laravel-ls/commit/f92475dcf64e067fe4baf9b57b72d87d9dbac8db))
- fix error checking bug. ([e82ffd6](https://github.com/laravel-ls/laravel-ls/commit/e82ffd687a3b6422a111b7d9386edc07acd9e409))
- adding editorconfig ([018f1c2](https://github.com/laravel-ls/laravel-ls/commit/018f1c2107b270bde4ab2da54a82dceea6c52257))
- fix bug in New() ([ef4b366](https://github.com/laravel-ls/laravel-ls/commit/ef4b366599023ded84eedd65a3adea86e29d7095))
- need to provide a MarshalJSON() for MarkupContentOrMarkedString ([62f3d69](https://github.com/laravel-ls/laravel-ls/commit/62f3d69c01487fb114abfcb3bf0720b77b200955))
- initialize protocol.HoverResult correctly. ([4435296](https://github.com/laravel-ls/laravel-ls/commit/44352963d95066a1e8318f7cdf25155c10be203f))
- set protocol.MarkupContent.Kind to markdown stead of markup ([e25090f](https://github.com/laravel-ls/laravel-ls/commit/e25090f4225d712168cf8f82b0105391ac795dc1))
- define MarkupKind type ([5e8351e](https://github.com/laravel-ls/laravel-ls/commit/5e8351e23bedb44ca84107c451ca91be364b5891))
- update Hover result to have protocol.MarkupKindMarkdown const instead of raw string. ([f94d99d](https://github.com/laravel-ls/laravel-ls/commit/f94d99d10daf16776288f3ee7d52244bd74acd95))
- fix documentation url ([718b4cc](https://github.com/laravel-ls/laravel-ls/commit/718b4cc7cd422c62e3fa54407cf4119caa3dff33))
- strings does not need to be pointers ([e5a8025](https://github.com/laravel-ls/laravel-ls/commit/e5a802563f69bcd7d940db13e376de77f5e58aa8))
- define MarshalJSON for DocumentDiagnosticReport ([b1c9afc](https://github.com/laravel-ls/laravel-ls/commit/b1c9afc3b22122dee2f1bbcbbd484aa38d2522b2))
- add logging for RPC messages ([25df126](https://github.com/laravel-ls/laravel-ls/commit/25df126adea6b96d788f76b55387048e2f62e44c))
- add support for windows ([0b4e750](https://github.com/laravel-ls/laravel-ls/commit/0b4e7501b0bfac4b3b1aab6389db165ae4f69e7f))
- use laravel-ls/uri package to correctly validate uri's ([98bec3a](https://github.com/laravel-ls/laravel-ls/commit/98bec3a5158a8475f98f3e5a272ee70d48a600ed))
- Adding lsp/protocol/document_change_operation.go ([85f0aaf](https://github.com/laravel-ls/laravel-ls/commit/85f0aafaab881765c03fd8c53f46a8165aa6a18f))
- change WorkspaceEdit.DocumentChanges type to []DocumentChangeOperation ([95b69d1](https://github.com/laravel-ls/laravel-ls/commit/95b69d1d5d28ffd0d22c1f8acda4a7e1819bbb16))
- use slices.Contains() ([0a87682](https://github.com/laravel-ls/laravel-ls/commit/0a87682a850f6059debc4eb88ae287b279796ae5))
- adding ResolveCodeAction ([11d9fc7](https://github.com/laravel-ls/laravel-ls/commit/11d9fc72ac26fbd24d8b2743fdbe165234adf245))
- adding new finder and view implementation ([3e09a91](https://github.com/laravel-ls/laravel-ls/commit/3e09a911346f26e2feab5c529e0df2db00ed580f))
- use new view finder implementation ([217d368](https://github.com/laravel-ls/laravel-ls/commit/217d368ee67244c1711f40ebe523e585e44ff7aa))
- Remove laravel/providers/view/filesystem.go (view finder is used instead) ([79c65e6](https://github.com/laravel-ls/laravel-ls/commit/79c65e6768b8016420029ba6d6d2ad0fc1352aaa))
- remove laravel/view_file.go (new implementation in laravel/view/view.go) ([86f2066](https://github.com/laravel-ls/laravel-ls/commit/86f20665aaeaa61f36f5ddb6ecf6c0bd3b6ab241))
- move code action to its own function ([53bd243](https://github.com/laravel-ls/laravel-ls/commit/53bd243f19d76cf329a29db8cc2864ca83a93e22))
- fix pointer values that do not need to be pointers. ([7aa74b6](https://github.com/laravel-ls/laravel-ls/commit/7aa74b6306c4cb220c33daf4c68fd5a6dc881ef9))
- in HandleTextDoucmentDidChange() need to pass the errs variable to errors.Join() otherwise it wil be overwritten. ([a382989](https://github.com/laravel-ls/laravel-ls/commit/a382989d8cb21c88562bf31aaaa122aef7efdaaa))
- pass root path to createViewCodeAction() ([47da030](https://github.com/laravel-ls/laravel-ls/commit/47da03037216ff5ecd7da7ea7740f8366b18e8ec))
- make Search() find files containing a given string, not just starting with. ([5cf7758](https://github.com/laravel-ls/laravel-ls/commit/5cf7758c3b1b062695b8d47afa86815e0d6d9159))
- refactor to laravel/asset/finder.go ([4e38fc4](https://github.com/laravel-ls/laravel-ls/commit/4e38fc45befd381859494b956921d3e42afa5bd7))
- go mod: laravel-ls and spf13/afero should not be indirect imported packages ([20455c6](https://github.com/laravel-ls/laravel-ls/commit/20455c66883f14d6f3be540bf708f2d02c85fee6))
- add error_reporting() to bootstrapSrc ([43c9b86](https://github.com/laravel-ls/laravel-ls/commit/43c9b8630fa164888bfd05f0e20f4c882f382739))

---
## [0.0.6](https://github.com/laravel-ls/laravel-ls/src/v0.0.6) - 2025-05-15

### Miscellaneous

- return ErrNoBinary variable instead of a new error. ([f03ead3](https://github.com/laravel-ls/laravel-ls/commit/f03ead3a11921b718ba5df91f214bd537017453c))
- rename to symbols.go ([dcf467b](https://github.com/laravel-ls/laravel-ls/commit/dcf467b5ca2d384186f7e21d54c0faf045642599))
- minor format fix in template ([7cb6211](https://github.com/laravel-ls/laravel-ls/commit/7cb621124589672d6fe4d8026cfcc1294ab5af5b))
- document download via github releases. ([a94cd69](https://github.com/laravel-ls/laravel-ls/commit/a94cd69c44b14fc0890cf95d305c471349696ce4))
- Add route section ([2749913](https://github.com/laravel-ls/laravel-ls/commit/27499138f133c1eb13142351351ba39af17da277))
- add query for response()->view() ([b97d726](https://github.com/laravel-ls/laravel-ls/commit/b97d7266fb0b9207030dbfafb5c103b37b03c085))

---
## [0.0.5](https://github.com/laravel-ls/laravel-ls/src/v0.0.5) - 2025-03-12

### Miscellaneous

- add viper ([2fc2871](https://github.com/laravel-ls/laravel-ls/commit/2fc2871370ce57d43a587e8503959f516846ee48))
- adding config struct ([5ce89ac](https://github.com/laravel-ls/laravel-ls/commit/5ce89ac326ce11df7576ed194be78d01500dfe32))
- use config and add more flags ([9c45a98](https://github.com/laravel-ls/laravel-ls/commit/9c45a98cf221c9cc4904f5d894bba4a6d9a2bf12))
- Adding config.example.yml ([a668e7b](https://github.com/laravel-ls/laravel-ls/commit/a668e7b7b4efa2194ca6d59437eff3ac913e2534))
- rename some context parameters to ctx to keep things consistent. ([bd19d66](https://github.com/laravel-ls/laravel-ls/commit/bd19d66da1265e172c7fa8336bfcb6fe75714e51))
- make sure we check the return value from viper.ReadInConfig() ([7fd9213](https://github.com/laravel-ls/laravel-ls/commit/7fd92136b3cd55eba1f2b24f9c7d51be514a0d61))
- linting fixes. ([1990de9](https://github.com/laravel-ls/laravel-ls/commit/1990de94aac9ffb76b98d824190dd92512d1afaf))
- go mod: cleanup ([b86afe2](https://github.com/laravel-ls/laravel-ls/commit/b86afe222ec5cb0aaa45e12523416f1ab85c11c2))
- remove debug logging statements ([5304655](https://github.com/laravel-ls/laravel-ls/commit/5304655082d8d11c635fb690e7d65621c4140826))
- use debug level for debug logging messages. ([69af4f9](https://github.com/laravel-ls/laravel-ls/commit/69af4f985e187c40318133f01e0625f553158dc6))

---
## [0.0.4](https://github.com/laravel-ls/laravel-ls/src/v0.0.4) - 2025-03-11

### Miscellaneous

- register for blade files also ([508bd04](https://github.com/laravel-ls/laravel-ls/commit/508bd04fea3f5434d759f5b030ebf76bdb128ed4))
- update queries to only fetch calls where the first argument is a raw string (no variables or concatinations) ([f8f2bfd](https://github.com/laravel-ls/laravel-ls/commit/f8f2bfda4bba8d0da6687ec8dc9eacc0706751f1))
- update queries to only fetch calls where the first argument is a raw string (no variables or concatinations) ([dbeaa7a](https://github.com/laravel-ls/laravel-ls/commit/dbeaa7a15eb4476094f1e4fe8fb9caecf3c2c601))
- cleanup dead code ([27788e4](https://github.com/laravel-ls/laravel-ls/commit/27788e49d42f1b00d413b2a6e2483dc0accd9c3a))
- code action: no need to advance to the next line when sending code action edits. ([6a1cbf7](https://github.com/laravel-ls/laravel-ls/commit/6a1cbf70d874f332f2af310abb0e92adc07d5b73))
- add getFile() helper that validates url and returns error if the file could not be opened ([b3e3d4b](https://github.com/laravel-ls/laravel-ls/commit/b3e3d4b09e6d658315eb07c70bd83d79551adad0))
- move injection queries from injections submodule to queries submodule. ([6c19a71](https://github.com/laravel-ls/laravel-ls/commit/6c19a71d62f6c90f5368ab689eff74838075a33a))
- add GetQuery() function. ([abba963](https://github.com/laravel-ls/laravel-ls/commit/abba963520a4cdb1537f5b2186d4c9dc7d07e2b2))
- move queries to treesitter/queries module ([af6e936](https://github.com/laravel-ls/laravel-ls/commit/af6e93621bd2d4dc5f12e1f14551bbf7ecfcc363))
- move .scm files to treesitter/assets package. ([09aec61](https://github.com/laravel-ls/laravel-ls/commit/09aec61d2293ca9c86288f4425d13f079de01860))
- move queries package into to treesitter package ([9a6d54b](https://github.com/laravel-ls/laravel-ls/commit/9a6d54b34b26730e06d8b56380dcb58529820375))
- add ErrLangNotSupported ([c871954](https://github.com/laravel-ls/laravel-ls/commit/c871954a925c1c568770566d2d7f21df2726e316))
- remove IsViewName() as it's not used. ([b35fcc0](https://github.com/laravel-ls/laravel-ls/commit/b35fcc0e664075acfd92fb1a9996ea22bd0222d9))
- Adding ReadQueryFromFile() and make GetQuery() and GetInjectionQuery() return a *ts.Query ([81828ef](https://github.com/laravel-ls/laravel-ls/commit/81828ef47556446c0c111a4489ecc6fcec4a652e))
- adding utils/cache package ([013c2ea](https://github.com/laravel-ls/laravel-ls/commit/013c2ea883f8f397920ea8d26de1ebd58127e12c))
- update linting rules and fix some naming issues. ([3041091](https://github.com/laravel-ls/laravel-ls/commit/304109170e0d33ace7ee22f3043e0350329bfc26))
- move info.go to program package ([5ec5f5d](https://github.com/laravel-ls/laravel-ls/commit/5ec5f5dcf63ceb05abba0a5c8dce4833cf4c6890))
- cache queries ([0da2225](https://github.com/laravel-ls/laravel-ls/commit/0da22252dbd11f94bad6adfcb0bbd709568c0b03))
- adding Forget() ([de54439](https://github.com/laravel-ls/laravel-ls/commit/de5443939dfc46a4adfaca1e127b12fca0be5d14))
- adding asset provider ([8fef69c](https://github.com/laravel-ls/laravel-ls/commit/8fef69cf4ca0c4a124e43f3ca014d20e6756986d))
- tick off assets features. ([12c6c73](https://github.com/laravel-ls/laravel-ls/commit/12c6c7305c561e9e82342a3be480d33ac41ac2a7))
- fix a bug where variable substitution outside of string was not taken into account ([bbba81f](https://github.com/laravel-ls/laravel-ls/commit/bbba81f1696031672cf93e589155be4278b7eb44))
- move rootCmd to cmd sub package ([1dbcdc1](https://github.com/laravel-ls/laravel-ls/commit/1dbcdc1f32895724e3a48ce7b6c205af94be0438))
- change Version to VersionOverride and provide a Version() function that fetches info from build if VersionOverride is empty. ([86055a1](https://github.com/laravel-ls/laravel-ls/commit/86055a14c28cb8d06e0b2ba69cb3dab2b4101b82))
- add version variable that forwards it to program.VersionOverride ([97c1281](https://github.com/laravel-ls/laravel-ls/commit/97c12814fd2bea9f2fdaf84f1072f5a3ace0c73a))
- Refactor treesitter language ([53d4f0e](https://github.com/laravel-ls/laravel-ls/commit/53d4f0e73614c7a048d5a68e10f664a54dcf64cd))
- refactor out the visualize code to its own file. ([a49cc6f](https://github.com/laravel-ls/laravel-ls/commit/a49cc6f8aeeacf411319b4f4ce8230c86129566d))
- linting fixes ([407ed42](https://github.com/laravel-ls/laravel-ls/commit/407ed4277b8f20208aadb4fd515bab9a144e5166))
- remove print statement ([1053583](https://github.com/laravel-ls/laravel-ls/commit/1053583aafe6c0e299c756605edd0294bbf39be6))
- update with installation documentation. ([5707812](https://github.com/laravel-ls/laravel-ls/commit/5707812d1cf2f370f2a98f46df691ed39b61a703))
- update install documentation ([0c89f80](https://github.com/laravel-ls/laravel-ls/commit/0c89f807890d49d6fc9919b0b77e1722f83faa59))
- Refactor treesitter language ([23e432f](https://github.com/laravel-ls/laravel-ls/commit/23e432fc0ce670fc8185309246ddadc21e5964fb))
- combine all queries into one: tags.scm ([5d6c6c8](https://github.com/laravel-ls/laravel-ls/commit/5d6c6c87b2740e81bef78a765878dcbd251a66dd))
- adding Name() ([6734a0e](https://github.com/laravel-ls/laravel-ls/commit/6734a0edab5e13f1eb9bfacadbfa46160b648698))
- implement capture cache ([01baf05](https://github.com/laravel-ls/laravel-ls/commit/01baf050fb04f830e8ec6d708df903fb4cd1c84c))
- Merge branch 'treesitter-tags' ([9a17353](https://github.com/laravel-ls/laravel-ls/commit/9a17353d7077ddf036456cbd34dc0c72c44fe9f1))
- Adding utils/file.go ([6219e31](https://github.com/laravel-ls/laravel-ls/commit/6219e31b034556ec0af678d506396539ec01b742))
- adding utils/repository ([ed30bd8](https://github.com/laravel-ls/laravel-ls/commit/ed30bd8c6257223199b884fa64e8ff0c33f77128))
- adding runtime package ([6286657](https://github.com/laravel-ls/laravel-ls/commit/628665711364267e745d12982c65bce9786271bc))
- move script things to project module ([85200dc](https://github.com/laravel-ls/laravel-ls/commit/85200dcaa81b68f8aaa56d4677249d5d1b060074))
- add generate ([875d905](https://github.com/laravel-ls/laravel-ls/commit/875d90553f97123b30521e6d8db2ac1bf1fdc462))
- use and initialize project module ([906e0fe](https://github.com/laravel-ls/laravel-ls/commit/906e0fec38ef7875ca35cd8e9bb4b93829946d45))
- add project ([c5ac4ce](https://github.com/laravel-ls/laravel-ls/commit/c5ac4ce4835fb64267dea1e7f8b8167d95a5bd87))
- pass project to each context ([2957a1f](https://github.com/laravel-ls/laravel-ls/commit/2957a1f048e8f580fa37bc34fb009536beb31533))
- Merge branch 'runtime' ([e0a2aac](https://github.com/laravel-ls/laravel-ls/commit/e0a2aac4be34658d34c7ee1189a5a1e06e60ebaa))
- adding app provider ([9e5931a](https://github.com/laravel-ls/laravel-ls/commit/9e5931a534c22a60c42109551f231f5765bded37))
- use app provider ([19c22e4](https://github.com/laravel-ls/laravel-ls/commit/19c22e4a88eba85629968cb21a10cc8289339f46))
- Merge branch 'app-provider' ([9cb0d49](https://github.com/laravel-ls/laravel-ls/commit/9cb0d49a3aa07d3799bab8e490e934df26385ee8))
- adding config provider ([1c28a24](https://github.com/laravel-ls/laravel-ls/commit/1c28a243fd24c0f5b4216bed7b0639cf96bbb351))
- use config provider. ([67b6fa0](https://github.com/laravel-ls/laravel-ls/commit/67b6fa041debfb0a7e2f7f3060d84ab787f1ac15))
- Merge branch 'config-provider' ([04e9f47](https://github.com/laravel-ls/laravel-ls/commit/04e9f47edaa2c1499832033f8792ca94a711a0d8))
- make Add() take a list ([70048e7](https://github.com/laravel-ls/laravel-ls/commit/70048e7fba924e48c792d342649910dc31f8fa54))
- add providers list to constructor ([2a4a703](https://github.com/laravel-ls/laravel-ls/commit/2a4a7037bf84670e9a819a54e8c098e663175767))
- pass all providers to provider.NewManager() ([0fd2eaf](https://github.com/laravel-ls/laravel-ls/commit/0fd2eaf16f9a0520aaa13eadd5edd28f4a3d28f9))
- add some more functions for app bindings ([4b499a9](https://github.com/laravel-ls/laravel-ls/commit/4b499a957ec8810b513f06c996c9e0290f95b2d3))
- document config and application bindings ([8051b61](https://github.com/laravel-ls/laravel-ls/commit/8051b619324332b07d5b4dfd703aa8644cfff6a2))
- linting ([dca5115](https://github.com/laravel-ls/laravel-ls/commit/dca51158d15c7c9a4c24405d8e14a6e1b366b77d))

### Refactoring

- refactor parser/env package to env/evaluator ([1c61463](https://github.com/laravel-ls/laravel-ls/commit/1c614639d13ee3f2dc47c6b7a508d3c863bb2cda))

---
## [0.0.3](https://github.com/laravel-ls/laravel-ls/src/v0.0.3) - 2025-01-07

### Miscellaneous

- update go modules ([c0e83d7](https://github.com/laravel-ls/laravel-ls/commit/c0e83d7ec978fd5b18f906634c77823d8a3c0ebf))
- Documentation update ([35c1112](https://github.com/laravel-ls/laravel-ls/commit/35c11123cebca5d046fa896f7c6b10d4bfa8e99d))
- rename RangeOverlap to RangeOverlayBytes and implement RangeOverlap with points. ([0d24a7c](https://github.com/laravel-ls/laravel-ls/commit/0d24a7ce308f1add811d49d42025740e4a57d492))
- Adding file/type_test.go ([b6bb16e](https://github.com/laravel-ls/laravel-ls/commit/b6bb16e10e66a64e30324e9df3d9cce43272c86a))
- properly handle env files. ([7b31a46](https://github.com/laravel-ls/laravel-ls/commit/7b31a46eb1eafd3c9031744e7959861a447b0dc6))
- adding structs related to code action ([89a279f](https://github.com/laravel-ls/laravel-ls/commit/89a279f326630983ac9e73aa2b644b63c1d81bf4))
- adding toTSRange() ([f0467e6](https://github.com/laravel-ls/laravel-ls/commit/f0467e6e06750acf16f1dafbc6ddb1f2a277bb87))
- Update readme ([66e82b6](https://github.com/laravel-ls/laravel-ls/commit/66e82b690656ef5afdbc7a9b75c81868f2769013))
- Update go.mod ([4a754c4](https://github.com/laravel-ls/laravel-ls/commit/4a754c4e50fc9ed7c2643b463eb070c3609ca851))
- changed to new module naming ([d36bcf1](https://github.com/laravel-ls/laravel-ls/commit/d36bcf1184c81171cf121c83619964c8a48f70d5))
- Update go.mod ([#1](https://github.com/laravel-ls/laravel-ls/issues/1)) ([9f53676](https://github.com/laravel-ls/laravel-ls/commit/9f53676b98da50e0f6554c355641a1d07f0d9a23))
- adding provider/code_action.go ([cc6dfe8](https://github.com/laravel-ls/laravel-ls/commit/cc6dfe8643bc1687ebbb15de44484386a60c606e))
- add support for code action providers ([6cab860](https://github.com/laravel-ls/laravel-ls/commit/6cab8607d4e1d9403b50019200771d2d6c869986))
- add code action handler ([ed17961](https://github.com/laravel-ls/laravel-ls/commit/ed179610d8e64928ca2e4cbfab19e6339c851407))
- adding EnvCallsInRange() ([9fd2038](https://github.com/laravel-ls/laravel-ls/commit/9fd20381a0dec003e72e00bc7467c15f8e3aae85))
- implement code action ([fe846e9](https://github.com/laravel-ls/laravel-ls/commit/fe846e9e772fccdc7fa782d19a391ba5b84a2c34))
- move Capture from file.go to capture.go ([4a2406d](https://github.com/laravel-ls/laravel-ls/commit/4a2406de99b71d0073d240d3932ea244b9032b2a))
- add CaptureSlice type and In() and At() methods ([1a5c452](https://github.com/laravel-ls/laravel-ls/commit/1a5c45247e0111d6517fb32bb46ac83e01832ea0))
- FindCaptures should return CaptureSlice ([b5736b0](https://github.com/laravel-ls/laravel-ls/commit/b5736b03e6dce858435b8b11f0cea7d73323a21a))
- Adding treesitter/php/helpers.go ([af8ec47](https://github.com/laravel-ls/laravel-ls/commit/af8ec479680f9a8e51f14df0b13da95a77b054fe))
- remove some query functions and use generic ones. ([979931c](https://github.com/laravel-ls/laravel-ls/commit/979931c9f49bff540ef0f8e45bc3a03d3f17740e))
- remove some query functions and use generic ones ([7d01593](https://github.com/laravel-ls/laravel-ls/commit/7d015936930b8f17493291a7d6e71cb1de4b9b4e))
- move to treesitter module ([8d3333b](https://github.com/laravel-ls/laravel-ls/commit/8d3333b0798a03e8eefdce78bf491597825c4e44))
- fix captial first letter in error string. ([9892777](https://github.com/laravel-ls/laravel-ls/commit/9892777d23281f324cbd0f11aa9b407c936968d8))
- Remove some debug logging statements ([ddfb487](https://github.com/laravel-ls/laravel-ls/commit/ddfb487f056f3f2110c3ae85e06499debddad1d2))
- mark code action for missing keys in .env files as complete ([f11950b](https://github.com/laravel-ls/laravel-ls/commit/f11950baa5f6929688af0c8ad28e716cfef4b742))
- update module name for version variable ([56c1d5c](https://github.com/laravel-ls/laravel-ls/commit/56c1d5cb6d8f086d630235db6cb82e2cce84e735))

---
## [0.0.2](https://github.com/laravel-ls/laravel-ls/src/v0.0.2) - 2025-01-03

### Bug Fixes

- fix linting issues ([b6edb16](https://github.com/laravel-ls/laravel-ls/commit/b6edb163b78814e9c3569f396bd88064719dfac1))

### Miscellaneous

- Create FUNDING.yml ([832f17e](https://github.com/laravel-ls/laravel-ls/commit/832f17e889d3931c605be2cc9474eb5a0259fe24))
- rename RangeIntersect to RangeOverlap ([8cc813d](https://github.com/laravel-ls/laravel-ls/commit/8cc813d76d7c7345743a70c45dae6d6a72983662))
- change IndentString to IndentSize ([f414ffa](https://github.com/laravel-ls/laravel-ls/commit/f414ffacd3e6234c225b2be3f6894f07846267ec))
- fix badge url ([2c4db71](https://github.com/laravel-ls/laravel-ls/commit/2c4db7163b319f6d3f49164f7220b0bae07e908c))
- use a simpler module name ([61a1d58](https://github.com/laravel-ls/laravel-ls/commit/61a1d5848a58c314e3ed77a2d4c01ffdb3789e1e))
- update readme ([ef0d598](https://github.com/laravel-ls/laravel-ls/commit/ef0d59860c610d0947a88b7b30546aa7c3a98789))
- align text ([03156b4](https://github.com/laravel-ls/laravel-ls/commit/03156b4dce2698f4652e2f44575433d4068a0a48))
- move test badge ([56fe13c](https://github.com/laravel-ls/laravel-ls/commit/56fe13cc8e636d6cb87f59c1026494d64d5069b9))
- Place test badge in the center ([f60ad07](https://github.com/laravel-ls/laravel-ls/commit/f60ad07655c5823de4675a014298de10417f867a))
- minor cleanup, use a lookup table instead of if statements in TypeByFilename() ([ea10c57](https://github.com/laravel-ls/laravel-ls/commit/ea10c57d55c2e34c2044ef031f633bada4f12039))
- remove parser/parsers.go as it is not used ([2ba6ce5](https://github.com/laravel-ls/laravel-ls/commit/2ba6ce5fe5457e31ddf2c290feae1ce575488c4d))
- adding .golangci.yml ([34f59b3](https://github.com/laravel-ls/laravel-ls/commit/34f59b35636280d2e5dbb62a9467ac2ca93b9928))
- errors not start with capital letters and not end with punctuation. ([b64df23](https://github.com/laravel-ls/laravel-ls/commit/b64df237cef47c2b59e423e3f3fc05f168c423fc))
- should not import logrus twice. ([2812a78](https://github.com/laravel-ls/laravel-ls/commit/2812a78d270a2472881ba5e188cdab7b76bc08e1))
- enable stylecheck ([13a7517](https://github.com/laravel-ls/laravel-ls/commit/13a7517fd4b0b6a4d1989355ce033ba3222ede65))
- error strings should start with lowercase words. ([84f5c48](https://github.com/laravel-ls/laravel-ls/commit/84f5c489ca703c9105b7d45dfc09235c5ff78422))
- use the same receiver name for all methods ([4b3487b](https://github.com/laravel-ls/laravel-ls/commit/4b3487b9ae8973429730614136479a9b295e3c69))
- use the same receiver name for all methods ([b105ce1](https://github.com/laravel-ls/laravel-ls/commit/b105ce13fdbe030a849ce80270c40d252daf9fb7))
- remove pointer receiver. ([b552f29](https://github.com/laravel-ls/laravel-ls/commit/b552f29577a4a6d1197e58d4e1d4d71e02dc585e))
- adding .github/workflows/lint.yml ([7ff9cd4](https://github.com/laravel-ls/laravel-ls/commit/7ff9cd4ae1aecba9c24bf2799191a886b026c348))
- add lint target ([b280f39](https://github.com/laravel-ls/laravel-ls/commit/b280f39167badb304efe8f97ab3c52c0e1603d3f))
- add test for PointInRange() and fix implementation ([6403081](https://github.com/laravel-ls/laravel-ls/commit/64030815426e336ec3a3e1dd3d9cc14415bcdb87))
- add test for RangeOverlap and fix implementation. ([75fbd63](https://github.com/laravel-ls/laravel-ls/commit/75fbd63f0e88d567d350905824b38fe8819ae15d))
- in GetViewName, also check if the node is a "encapsed_string" node. ([bf860b3](https://github.com/laravel-ls/laravel-ls/commit/bf860b3c91060d7819e41ac9f4a4c92cdfa4cdf2))
- minor comment update ([43a7633](https://github.com/laravel-ls/laravel-ls/commit/43a7633041ec1dfdf3080ffeefc241b93e147dd4))
- in HandleTextDocumentHover() add file cache in hover context ([39b67e3](https://github.com/laravel-ls/laravel-ls/commit/39b67e3d9a529c1b0e04a6554914942da6e5bbde))
- in HandleTextDocumentCompletion() add file cache to completion context ([161540a](https://github.com/laravel-ls/laravel-ls/commit/161540af5169c68ee8ad504e68658994d07a1e8d))
- follow the other provider patterns by using a publish callback in the context instead of return value ([19adcac](https://github.com/laravel-ls/laravel-ls/commit/19adcac96ff8986d290b91ade0f854a7a9e34940))
- follow the other provider patterns by using a publish callback in the context instead of return value ([a3d9cab](https://github.com/laravel-ls/laravel-ls/commit/a3d9cab248bf44a91d924067b2aa70822aa4b4f1))
- in HandleTextDocumentDiagnostic() add file cache to context ([72660c2](https://github.com/laravel-ls/laravel-ls/commit/72660c2e94adf8595ae7bed79140073045bac6ae))
- remove pointer receiver for method that dont need it. ([80e8a82](https://github.com/laravel-ls/laravel-ls/commit/80e8a826ceda057c680d9fa8f752a4b2c45b2349))
- no need to pass child trees to newLanguageTree() ([7ed00a4](https://github.com/laravel-ls/laravel-ls/commit/7ed00a4f03989c6713eab74e56451ddc4f94f3cd))
- Hover() must not return a string ([1cba35f](https://github.com/laravel-ls/laravel-ls/commit/1cba35f5e1db4f55902e3c79adfaa308296b8827))
- add more tests for PointInRange ([4b48e2f](https://github.com/laravel-ls/laravel-ls/commit/4b48e2f75227da0e125a6d47e9d4792c91199a62))
- implement PointInRange correctly (for the second time :P) ([abf7af2](https://github.com/laravel-ls/laravel-ls/commit/abf7af2589c09d0b3f080768f5acb8927baa363b))
- do not populate the repositor on init. do it before querying the repo. ([d731503](https://github.com/laravel-ls/laravel-ls/commit/d7315032fd7fe7bb4d806e811b467fa5bb68cb11))
- in Hover() show "undefined" for keys that are undefined. ([ef8c373](https://github.com/laravel-ls/laravel-ls/commit/ef8c373c6080e032ba29fe7e4542718a7abc688b))
- add HasDefault() ([3372e29](https://github.com/laravel-ls/laravel-ls/commit/3372e299fd5d47eeac669f69a616735727b29331))
- only show diagnostics for keys that does not have default values. ([9f5b7a5](https://github.com/laravel-ls/laravel-ls/commit/9f5b7a53616d43f1b7f135fdc6e186a5e0e67cf0))

### Style

- style fixes for imports ([d88fc90](https://github.com/laravel-ls/laravel-ls/commit/d88fc90a80e4fd519f818b939963dc3205951dd2))

---
## [0.0.1](https://github.com/laravel-ls/laravel-ls/src/v0.0.1) - 2024-12-29

### Miscellaneous

- Initial commit ([f38b20c](https://github.com/laravel-ls/laravel-ls/commit/f38b20c314ac6fa43e8f3c3949ff0d264c797dc0))
- fix operation precendence in if statement ([af94a21](https://github.com/laravel-ls/laravel-ls/commit/af94a21dc914332a17983ad24d807da70e289c3c))
- pass args to binary ([ae5e99b](https://github.com/laravel-ls/laravel-ls/commit/ae5e99bee8587e2cf2b2a255a683af2c7007db5c))
- adding info.go ([61715a7](https://github.com/laravel-ls/laravel-ls/commit/61715a7e7ba120702ce5f8e4e9c5307df68bb3bd))
- use info.go to get program name and version ([8280106](https://github.com/laravel-ls/laravel-ls/commit/82801065f60d70b645f813e6f7a0993341956093))
- adding build flags ([cb17ec4](https://github.com/laravel-ls/laravel-ls/commit/cb17ec4466f5f38edb335171059eee8757db7f65))
- populate Version in info.go with info from git. ([6fe26fd](https://github.com/laravel-ls/laravel-ls/commit/6fe26fdc55d62898c50afce6b26ec3f647efd83c))
- Adding .github/workflows/test.yml ([063fe6c](https://github.com/laravel-ls/laravel-ls/commit/063fe6ce2ed37b43c5a427079b5f4d9c744207b5))
- Adding .github/workflows/release.yml ([3543ee1](https://github.com/laravel-ls/laravel-ls/commit/3543ee1c6ac4e4ceec1a9a6d77eca150ac772d52))
- only compile for native amd64 ([44a1474](https://github.com/laravel-ls/laravel-ls/commit/44a1474e8cf14a4ad55157c554e83456016e64cc))
- run on ubuntu-24.04 ([745e245](https://github.com/laravel-ls/laravel-ls/commit/745e2451b6c824633bbd5fc89ce5c4be5e83ea45))
- update ([f2bb549](https://github.com/laravel-ls/laravel-ls/commit/f2bb54944ff0db911b233de7f98949c52841e8d5))

<!-- generated by git-cliff -->
