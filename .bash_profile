# Avoid duplicate entries in .bash_history
export HISTCONTROL=ignoreboth:erasedups

# Cleaner terminal "$ prompt"
# https://gist.github.com/vargeorge/8b20488b7d7b6c101b4b
export PS1="\W $ "

# Extract GitHub token from Git credential storage
extract_github_token() {
  echo "url=https://github.com" | git credential fill | grep "^password=" | cut -d= -f2;
}
