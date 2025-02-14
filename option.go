package gdrive // nolint: golint

// Option can be used to pass optional Options to GDriver
type Option func(driver *GDriver) error

// RootDirectory sets the root directory for all operations
func RootDirectory(path string) Option {
	return func(driver *GDriver) error {
		_, err := driver.SetRootDirectory(path)
		return err
	}
}

// RootNodeId sets the root directory for all operations
func RootNode(id string) Option {
	return func(driver *GDriver) error {
		err := driver.SetRootNode(id)
		if err != nil {
			return err
		}
		_, err = driver.SetRootDirectory("")
		return err
	}
}