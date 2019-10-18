package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"./model"
	"log"
	"time"
)

var eg *xorm.EngineGroup
func init(){
	log.SetFlags(log.LstdFlags|log.Lshortfile)
}
func main() {
	conns := []string{
		"root:root@(localhost:3306)/test",
		"root:root@(localhost:3307)/test",
	}

	var err error
	eg, err = xorm.NewEngineGroup("mysql", conns)
	if err != nil {
		log.Println(err)
		return
	}
	user := new(model.Users)
	user.Name = "myname"
	user.PrefList="{}"
	user.LatestLogTime=time.Now()
	affected, err := eg.Insert(user)
	// INSERT INTO user (name) values (?)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(affected)
}
/*
CREATE TABLE `news` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `content` text,
  `news_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `title` text NOT NULL,
  `module_id` int(11) NOT NULL,
  `url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=332 DEFAULT CHARSET=utf8;
*/