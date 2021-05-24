package cmd

import (
	"bytes"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	completionCmd.InitDefaultHelpFlag()
}

var completionCmd = &cobra.Command{
	Use:               "completion",
	Short:             "generates zsh completions",
	SilenceErrors:     true,
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error { return nil },
	RunE: func(cmd *cobra.Command, _ []string) error {
		// code copied from https://github.com/kubernetes/kubectl/blob/031e3ab456e6d0e676aecd163e52fcb4906e1faf/pkg/cmd/completion/completion.go

		f, err := os.Create("_procgo")
		if err != nil {
			return err
		}

		zshHead := "#compdef procgo\n"
		f.Write([]byte(zshHead))

		zshInitialization := `
__procgo_bash_source() {
	alias shopt=':'
	emulate -L sh
	setopt kshglob noshglob braceexpand
	source "$@"
}
__procgo_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift
		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on
		# all the time
		if [ "$1" = "__procgo_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}
__procgo_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?
	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}
__procgo_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}
__procgo_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}
__procgo_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}
__procgo_filedir() {
	# Don't need to do anything here.
	# Otherwise we will get trailing space without "compopt -o nospace"
	true
}
autoload -U +X bashcompinit && bashcompinit
# use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --version 2>&1 | grep -q GNU; then
	LWORD='\<'
	RWORD='\>'
fi
__procgo_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__procgo_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__procgo_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__procgo_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__procgo_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__procgo_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__procgo_type/g" \
	<<'BASH_COMPLETION_EOF'
`
		f.Write([]byte(zshInitialization))

		buf := new(bytes.Buffer)
		cmd.Root().GenBashCompletion(buf)
		f.Write(buf.Bytes())

		zshTail := `
BASH_COMPLETION_EOF
}
__procgo_bash_source <(__procgo_convert_bash_to_zsh)
`
		f.Write([]byte(zshTail))
		return nil
	},
}
