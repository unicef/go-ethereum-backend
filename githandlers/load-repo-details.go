package githandlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/qjouda/dignity-platform/backend/apperrors"
	"github.com/qjouda/dignity-platform/backend/datatype"
	"gopkg.in/src-d/go-git.v4"
)

func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadRepoDetails loads details of the passed Git link (project name, description, branches)
// @TODO integration with git is not complete due to project time termincation
func LoadRepoDetails(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := struct {
			URL string `json:"url"`
		}{}
		c.BindJSON(&body)
		path := "tmp/repo"
		err := removeContents(path)
		if err != nil {
			apperrors.Critical("githandlers/load-repo-details:0", err)
			c.String(http.StatusServiceUnavailable, datatype.ErrServerError.Error())
		}
		repo, err := git.PlainClone(path, false, &git.CloneOptions{URL: body.URL})
		if err != nil {
			apperrors.Critical("githandlers/load-repo-details:1", err)
			c.String(http.StatusBadRequest, err.Error())
		}
		ref, err := repo.Head()
		if err != nil {
			apperrors.Critical("githandlers/load-repo-details:2", err)
			c.String(http.StatusBadRequest, err.Error())
		}
		log.Println("HEAD ref is >>> ", ref.Hash())
		commit, err := repo.CommitObject(ref.Hash())
		log.Println("Last commit is >>>>", commit)
		c.String(http.StatusOK, ref.Hash().String())
	}
}
