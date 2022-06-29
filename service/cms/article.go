package cms

import (
	"log"
	"strconv"
	"strings"
	"wave-admin/global"
	"wave-admin/model/cms"
	"wave-admin/model/cms/request"
	"wave-admin/model/cms/response"
)

type ArticleService struct{}

func (articleService *ArticleService) GetArticle(id int) (err error, repArticle response.Article) {
	var article cms.CmsArticle
	err = global.GnDb.Where("id = ?", id).Find(&article).Error

	var articleTags []cms.CmsArticleTag
	err = global.GnDb.Where("article_id = ?", id).Find(&articleTags).Error

	var articleContent cms.CmsArticleContent
	err = global.GnDb.Where("article_id = ?", id).Find(&articleContent).Error

	var tags []int
	for _, tag := range articleTags {
		tags = append(tags, int(tag.TagId))
	}
	repArticle = response.Article{
		ID:          article.ID,
		Status:      strconv.Itoa(int(article.Status)),
		Title:       article.Title,
		Img:         article.Img,
		KeyWord:     article.KeyWord,
		CategoryId:  article.CategoryId,
		Description: articleContent.Description,
		Content:     articleContent.Content,
		Tag:         tags,
	}

	return err, repArticle
}

func (articleService *ArticleService) AddArticle(userId uint64, article request.ReqArticle) (err error) {
	add := cms.CmsArticle{
		Status:      article.Status,
		Title:       article.Title,
		UserID:      userId,
		EditId:      userId,
		Img:         article.Img,
		KeyWord:     article.KeyWord,
		CategoryId:  article.CategoryId,
	}
	err = global.GnDb.Create(&add).Error
	if len(article.TagIds) > 0 {
		err = articleService.addArticleTag(add.ID, article.TagIds)
	}
	err = articleService.addArticleContent(add.ID, article)

	return err
}

func (articleService *ArticleService) addArticleTag(articleId uint64, tagIds string) (err error) {
	global.GnDb.Where("article_id = ?", articleId).Unscoped().Delete(&cms.CmsArticleTag{})

	ids := strings.Split(tagIds, ",")

	var articleTags []cms.CmsArticleTag
	for _, tagId := range ids {
		id,_ := strconv.Atoi(tagId)
		articleTags = append(articleTags, cms.CmsArticleTag{
			ArticleId: articleId,
			TagId: uint64(id),
		})
	}
	err = global.GnDb.Model(&cms.CmsArticleTag{}).Create(&articleTags).Error

	return err
}

func (articleService *ArticleService) addArticleContent(id uint64, article request.ReqArticle) (err error) {
	var count int64
	global.GnDb.Model(&cms.CmsArticleContent{}).Where("article_id = ?", id).Count(&count)
	if count > 0 {
		update := cms.CmsArticleContent{
			Description: article.Description,
			Content:     article.Content,
		}
		err = global.GnDb.Model(&cms.CmsArticleContent{}).Where("article_id = ?", id).Updates(update).Error
	} else {
		content := cms.CmsArticleContent{
			ArticleId:   id,
			Description: article.Description,
			Content:     article.Content,
		}
		err = global.GnDb.Create(&content).Error
	}

	return err
}

func (articleService *ArticleService) UpdateArticle(userId uint64, article request.ReqArticle, id int) (err error) {
	update := cms.CmsArticle{
		Status:      article.Status,
		Title:       article.Title,
		EditId:      userId,
		Img:         article.Img,
		KeyWord:     article.KeyWord,
		CategoryId:  article.CategoryId,
	}
	err = global.GnDb.Model(&cms.CmsArticle{}).Where("id = ?", id).Updates(update).Error
	log.Println("tags=", article.TagIds)
	if len(article.TagIds) > 0 {
		err = articleService.addArticleTag(uint64(id), article.TagIds)
	}
	err = articleService.addArticleContent(uint64(id), article)

	return err
}

func (articleService *ArticleService) DeleteArticle(id int) (err error) {
	err = global.GnDb.Delete(&cms.CmsArticleTag{}, "article_id = ?", id).Error
	if err != nil {
		return err
	}
	err = global.GnDb.Delete(&cms.CmsArticleContent{}, "article_id = ?", id).Error
	if err != nil {
		return err
	}
	err = global.GnDb.Delete(&cms.CmsArticle{}, "id = ?", id).Error

	return err
}

func (articleService *ArticleService) GetArticleList(info request.ArticleList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.CurrentPage - 1)
	var articles []cms.CmsArticle
	db := global.GnDb.Model(&cms.CmsArticle{}).Where("status = ?", info.Status)
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, articles, total
	}

	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&articles).Error

	return err, articles, total
}
