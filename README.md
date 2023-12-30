# Various Tools

## GitHub Best Practices

1. Branch naming conventions often include:
   - `feature/`: For new features (e.g. feature/add-login)
   - `bugfix/`: For bug fixes (e.g. bugfix/fix-login-error)
   - `hotfix/`: For urgent fixes (e.g. hotfix/patch-security-issue)

---

## Dev Tools for MacOS/Linux

1. [op-secrets](./op-secrets/) CLI
2. `.bash_profile`
3. `.npmrc`

#### General

1. [Git Credential Manager](https://github.com/GitCredentialManager/git-credential-manager)
   - Use macOS Keychain to store creds `$ git config --global credential.helper osxkeychain`
2. [1Password CLI](https://developer.1password.com/docs/cli/get-started)
   - `$ brew install --cask 1password/tap/1password-cli`
   - `$ op signin`
3. [Docker Desktop](https://www.docker.com/products/docker-desktop)
4. [Task Runner v3.3x](https://taskfile.dev/installation) (simpler Make alternative)

#### NodeJS

1. [NodeJS v20](https://nodejs.org)
2. [Pnpm Package Manager](https://pnpm.io)

#### Web / Sveltekit

1. [Chrome DevTools](https://github.com/sveltejs/svelte-devtools)

#### Python

1. [Python 3.9](https://docs.python-guide.org/starting/install3/osx)
2. [Poetry Package Manager](https://python-poetry.org)
3. [Poetry Bash, Fish, or Zsh Completion](https://python-poetry.org)

---

## VSCode Plugins

You can copy the `./vscode` directory to your workspace to install all the VsCode plugins and settings.

#### General

1. Markdown (https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one)
2. Prettier (https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)
3. Bracket Pair Colorization (https://marketplace.visualstudio.com/items?itemName=dzhavat.bracket-pair-toggler)
4. DotENV (https://marketplace.visualstudio.com/items?itemName=mikestead.dotenv)
5. Better Comments (https://marketplace.visualstudio.com/items?itemName=aaron-bond.better-comments)
6. Error Lens (https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens)
7. Even Better TOML (https://marketplace.visualstudio.com/items?itemName=tamasfe.even-better-toml)
8. YAML (https://marketplace.visualstudio.com/items?itemName=redhat.vscode-yaml)
9. Change Case (https://marketplace.visualstudio.com/items?itemName=wmaurer.change-case)

#### NodeJS

1. ESLint (https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
2. Pretty TypeScript Errors (https://marketplace.visualstudio.com/items?itemName=yoavbls.pretty-ts-errors)
3. Node Modules Intellisense (https://marketplace.visualstudio.com/items?itemName=leizongmin.node-module-intellisense)

#### Web / Sveltekit

1. Svelte (https://marketplace.visualstudio.com/items?itemName=svelte.svelte-vscode)
2. Svelte Intellisense (https://marketplace.visualstudio.com/items?itemName=ardenivanov.svelte-intellisense)
3. Tailwind IntelliSense (https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
4. PostCSS Support (https://marketplace.visualstudio.com/items?itemName=csstools.postcss)
5. SVG (https://marketplace.visualstudio.com/items?itemName=jock.svg)
6. Auto Rename Tag (https://marketplace.visualstudio.com/items?itemName=formulahendry.auto-rename-tag)
7. Auto Close Tag (https://marketplace.visualstudio.com/items?itemName=formulahendry.auto-close-tag)
8. HTML CSS Support (https://marketplace.visualstudio.com/items?itemName=ecmel.vscode-html-css)
9. CSS Peek (https://marketplace.visualstudio.com/items?itemName=pranaygp.vscode-css-peek)

#### Python

1. Python (https://marketplace.visualstudio.com/items?itemName=ms-python.python)
2. Pylance (https://marketplace.visualstudio.com/items?itemName=ms-python.vscode-pylance)

---

## Misc CLI Tools

1. https://github.com/cube2222/octosql - `OctoSQL` is a query tool that allows you to join, analyse and transform data from multiple databases and file formats using SQL
2. https://github.com/sharkdp/bat - Better bat with syntax highlighting
3. https://github.com/rs/curlie - Better curl
4. https://github.com/dbcli/pgcli - Postgres CLI
