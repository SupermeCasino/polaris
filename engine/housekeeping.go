package engine

import (
	"polaris/ent/media"
	"polaris/log"
)

func (c *Engine) housekeeping() error {
	log.Infof("start housekeeping tasks...")

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
	count := 0
	log.Infof("check db scraps, total episodes: %v, total media: %v", len(allEpisodes), len(validMediaIDs))
	for _, ep := range allEpisodes {
		if _, ok := validMediaIDs[ep.MediaID]; !ok {
			//log.Infof("remove scrap episode record: %v S%vE%v", ep.MediaID, ep.SeasonNumber, ep.EpisodeNumber)
			if err := c.db.DeleteEpisode(ep.ID); err != nil {
				log.Errorf("delete scrap episode record error: %v", err)
			}
			count++
			if count%100 == 0 {
				log.Infof("%v scrap episode records removed...", count)
			}
		}
	}
	log.Infof("%v scrap episode records removed...", count)

	return nil
}

func (c *Engine) checkImageFilesInterity() error {
	//TODO: download missing image files, remove unused image files
	return nil
}
