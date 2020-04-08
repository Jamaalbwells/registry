package models

import (
	"context"
	"fmt"
	"regexp"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

// We might extend this to all characters that do not require escaping.
// See "Resource ID Segments" in https://aip.dev/122.
const nameRegex = "([a-zA-Z0-9-_\\.]+)"

// File names include all of the above characters plus forward slashes.
const fileNameRegex = "([a-zA-Z0-9-_\\.\\/]+)"

func validateID(id string) error {
	r := regexp.MustCompile("^" + nameRegex + "$")
	m := r.FindAllStringSubmatch(id, -1)
	if m == nil {
		return fmt.Errorf("invalid id '%s'", id)
	}
	return nil
}

func validateFileID(id string) error {
	r := regexp.MustCompile("^" + fileNameRegex + "$")
	m := r.FindAllStringSubmatch(id, -1)
	if m == nil {
		return fmt.Errorf("invalid id '%s'", id)
	}
	return nil
}

// ProductsRegexp returns a regular expression that matches collection of products.
func ProductsRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products$")
}

// ProductRegexp returns a regular expression that matches a product resource name.
func ProductRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "$")
}

// VersionsRegexp returns a regular expression that matches a collection of versions.
func VersionsRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "/versions$")
}

// VersionRegexp returns a regular expression that matches a version resource name.
func VersionRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "/versions/" + nameRegex + "$")
}

// SpecsRegexp returns a regular expression that matches a collection of specs.
func SpecsRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "/versions/" + nameRegex + "/specs$")
}

// SpecRegexp returns a regular expression that matches a spec resource name.
func SpecRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "/versions/" + nameRegex + "/specs/" + nameRegex + "$")
}

// FilesRegexp returns a regular expression that matches a collection of files.
func FilesRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "/versions/" + nameRegex + "/specs/" + nameRegex + "/files$")
}

// FileRegexp returns a regular expression that matches a file resource name.
func FileRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex + "/products/" + nameRegex + "/versions/" + nameRegex + "/specs/" + nameRegex + "/files/" + fileNameRegex + "$")
}

func deleteAllMatches(ctx context.Context, client *datastore.Client, q *datastore.Query) error {
	it := client.Run(ctx, q.Distinct())
	key, err := it.Next(nil)
	keys := make([]*datastore.Key, 0)
	for err == nil {
		keys = append(keys, key)
		key, err = it.Next(nil)
	}
	if err != iterator.Done {
		return err
	}
	return client.DeleteMulti(ctx, keys)
}
