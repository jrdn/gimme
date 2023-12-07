package source

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/j13g/goutil/errs"
	"github.com/j13g/goutil/regex"
)

var (
	githubRegex = regexp.MustCompile(`^(?P<owner>.*)/(?P<repo>.*)@(?P<version>.*)$`)
	gitUrl      = regexp.MustCompile(`^(?P<url>(?:https|git).*\.git)(@(?P<version>.*))?$`)
)

var ErrParseSourceFailed = errors.New("failed to parse source")

func Parse(s string) (Source, error) {
	if matches := regex.Match(gitUrl, s); len(matches) > 0 {
		return parseGitSpec(matches[0])
	}

	if matches := regex.Match(githubRegex, s); len(matches) > 0 {
		return parseGithubSpec(matches[0])
	}

	return nil, ErrParseSourceFailed
}

func parseGitSpec(match map[string]string) (Source, error) {
	var url, version string
	var urlOk bool

	if url, urlOk = match["url"]; !urlOk || url == "" {
		return nil, errs.Wrap(ErrParseSourceFailed, "git source missing url")
	}

	version = match["version"]
	return &gitSource{
		pullURL: url,
		version: version,
	}, nil
}

func parseGithubSpec(match map[string]string) (Source, error) {
	var owner, repo, version string
	var ownerOk, repoOk, versionOk bool

	if owner, ownerOk = match["owner"]; !ownerOk || owner == "" {
		return nil, errs.Wrap(ErrParseSourceFailed, "github source missing owner")
	}

	if repo, repoOk = match["repo"]; !repoOk || repo == "" {
		return nil, errs.Wrap(ErrParseSourceFailed, "github source missing repo")
	}

	if version, versionOk = match["version"]; !versionOk || version == "" {
		return nil, errs.Wrap(ErrParseSourceFailed, "github source missing version")
	}

	return &gitSource{
		pullURL: fmt.Sprintf("https://github.com/%s/%s.git", owner, repo),
		version: version,
	}, nil
}
