package blog

import dbmysql "git.services.wait/chenwx/cwxgoweb/src/common/dbmysql"

type docDB struct {
	db *dbmysql.MysqlConf
}

func (d *docDB) GetDocPv(url string) int {
	session := d.db.GetSession()
	var doc T_doc_access

	return doc.getPv(session, url)
}
