package engine

import (
	"polaris/ent/media"
	"polaris/log"
)

func (c *Engine) housekeeping() (err error) {
	log.Infof("start housekeeping tasks...")
	defer func() {
		log.Infof("housekeeping tasks completed. err: %v", err)
	}()

	if err := c.checkDbScraps(); err != nil {
		return err
	}
	if err := c.checkImageFilesInterity(); err != nil {
		return err
	}

	return nil
}

func (c *Engine) checkDbScraps() error {
	//TODO: remove episodes that are not associated with any series

	tvs := c.db.GetMediaWatchlist(media.MediaTypeTv)
	movies := c.db.GetMediaWatchlist(media.MediaTypeMovie)

	validMediaIDs := make(map[int]bool, len(tvs)+len(movies))
	for _, tv := range tvs {
		validMediaIDs[tv.ID] = true
	}
	for _, movie := range movies {
		validMediaIDs[movie.ID] = true
	}

	allEpisodes, err := c.db.GetAllEpisodes()
	if err != nil {
		log.Debugf("get all episodes error: %v", err)
		return err
	}

	log.Infof("check db scraps, total episodes: %v, total media: %v", len(allEpisodes), len(validMediaIDs))
	toDeleteIds := make([]int, 0)
	for _, ep := range allEpisodes {
		if _, ok := validMediaIDs[ep.MediaID]; !ok {
			//log.Infof("remove scrap episode record: %v S%vE%v", ep.MediaID, ep.SeasonNumber, ep.EpisodeNumber)
			toDeleteIds = append(toDeleteIds, ep.ID)
		}
	}

	log.Infof("%v scrap episode records will be removed...", len(toDeleteIds))

	if err := c.db.DeleteEpisode(toDeleteIds...); err != nil {
		log.Errorf("delete scrap episode records error: %v", err)
	}

	return nil
}

func (c *Engine) checkImageFilesInterity() error {
	//TODO: download missing image files, remove unused image files
	return nil
}
