package models

type History struct {
	Name string `bson:"name"`
	UID  string `bson:"uid"`
	Text string `bson:"text"`
	CID  string `bson:"id"`
	MID  string `bson:"id"`
}
