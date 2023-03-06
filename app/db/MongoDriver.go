package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ObjectIDFromHex = func(hex string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(hex)

	if err != nil {
		log.Fatal(err)
	}
	return objectID
}
var IsObjectId = func(id string) bool {
	return primitive.IsValidObjectID(id)
}

// 各个表的Collection对象
var Notebooks *mongo.Collection
var Notes *mongo.Collection
var NoteContents *mongo.Collection
var NoteContentHistories *mongo.Collection

var ShareNotes *mongo.Collection
var ShareNotebooks *mongo.Collection
var HasShareNotes *mongo.Collection
var Blogs *mongo.Collection
var Users *mongo.Collection
var Groups *mongo.Collection
var GroupUsers *mongo.Collection

var Tags *mongo.Collection
var NoteTags *mongo.Collection
var TagCounts *mongo.Collection

var UserBlogs *mongo.Collection

var Tokens *mongo.Collection

var Suggestions *mongo.Collection

// Album & file(image)
var Albums *mongo.Collection
var Files *mongo.Collection
var Attachs *mongo.Collection

var NoteImages *mongo.Collection
var Configs *mongo.Collection
var EmailLogs *mongo.Collection

// blog
var BlogLikes *mongo.Collection
var BlogComments *mongo.Collection
var Reports *mongo.Collection
var BlogSingles *mongo.Collection
var Themes *mongo.Collection
var Sessions *mongo.Collection

func init() {
	initMongo("mongodb://localhost:11003", "leanote")
}

func initMongo(url, dbname string) {
	// 不关闭 ?
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	// notebook
	Notebooks = client.Database(dbname).Collection("notebooks")

	// notes
	Notes = client.Database(dbname).Collection("notes")

	// noteContents
	NoteContents = client.Database(dbname).Collection("note_contents")
	NoteContentHistories = client.Database(dbname).Collection("note_content_histories")

	// share
	ShareNotes = client.Database(dbname).Collection("share_notes")
	ShareNotebooks = client.Database(dbname).Collection("share_notebooks")
	HasShareNotes = client.Database(dbname).Collection("has_share_notes")

	// user
	Users = client.Database(dbname).Collection("users")
	// group
	Groups = client.Database(dbname).Collection("groups")
	GroupUsers = client.Database(dbname).Collection("group_users")

	// blog
	Blogs = client.Database(dbname).Collection("blogs")

	// tag
	Tags = client.Database(dbname).Collection("tags")
	NoteTags = client.Database(dbname).Collection("note_tags")
	TagCounts = client.Database(dbname).Collection("tag_count")

	// blog
	UserBlogs = client.Database(dbname).Collection("user_blogs")
	BlogSingles = client.Database(dbname).Collection("blog_singles")
	Themes = client.Database(dbname).Collection("themes")

	// find password
	Tokens = client.Database(dbname).Collection("tokens")

	// Suggestion
	Suggestions = client.Database(dbname).Collection("suggestions")

	// Album & file
	Albums = client.Database(dbname).Collection("albums")
	Files = client.Database(dbname).Collection("files")
	Attachs = client.Database(dbname).Collection("attachs")

	NoteImages = client.Database(dbname).Collection("note_images")

	Configs = client.Database(dbname).Collection("configs")
	EmailLogs = client.Database(dbname).Collection("email_logs")

	// 社交
	BlogLikes = client.Database(dbname).Collection("blog_likes")
	BlogComments = client.Database(dbname).Collection("blog_comments")

	// 举报
	Reports = client.Database(dbname).Collection("reports")

	// session
	Sessions = client.Database(dbname).Collection("sessions")

}

func close() {

	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}

}

// common DAO
// 公用方法

//----------------------

func Insert(collection *mongo.Collection, i interface{}) bool {
	_, err := collection.InsertOne(context.Background(), i)
	return Err(err)
}

//----------------------

// 适合一条记录全部更新
func Update(collection *mongo.Collection, query interface{}, i interface{}) bool {
	_, err := collection.UpdateOne(context.Background(), query, i)
	return Err(err)
}

func Upsert(collection *mongo.Collection, query interface{}, i interface{}) bool {
	upsert := true
	options := options.UpdateOptions{Upsert: &upsert}
	_, err := collection.UpdateOne(context.Background(), query, i, &options)
	return Err(err)
}
func UpdateAll(collection *mongo.Collection, query interface{}, i interface{}) bool {
	_, err := collection.UpdateMany(context.Background(), query, i)
	return Err(err)
}
func UpdateByIdAndUserId(collection *mongo.Collection, id, userId string, i interface{}) bool {
	_, err := collection.UpdateOne(context.Background(), GetIdAndUserIdQ(id, userId), i)
	return Err(err)
}

func UpdateByIdAndUserId2(collection *mongo.Collection, id, userId string, i interface{}) bool {
	_, err := collection.UpdateOne(context.Background(), GetIdAndUserIdBsonQ(id, userId), i)
	return Err(err)
}
func UpdateByIdAndUserIdField(collection *mongo.Collection, id, userId, field string, value interface{}) bool {
	return UpdateByIdAndUserId(collection, id, userId, bson.M{"$set": bson.M{field: value}})
}
func UpdateByIdAndUserIdMap(collection *mongo.Collection, id, userId string, v bson.M) bool {
	return UpdateByIdAndUserId(collection, id, userId, bson.M{"$set": v})
}

func UpdateByIdAndUserIdField2(collection *mongo.Collection, id, userId string, field string, value interface{}) bool {
	return UpdateByIdAndUserId2(collection, id, userId, bson.M{"$set": bson.M{field: value}})
}
func UpdateByIdAndUserIdMap2(collection *mongo.Collection, id, userId string, v bson.M) bool {
	return UpdateByIdAndUserId2(collection, id, userId, bson.M{"$set": v})
}

func UpdateByQField(collection *mongo.Collection, q interface{}, field string, value interface{}) bool {
	_, err := collection.UpdateMany(context.Background(), q, bson.M{"$set": bson.M{field: value}})
	return Err(err)
}
func UpdateByQI(collection *mongo.Collection, q interface{}, v interface{}) bool {
	_, err := collection.UpdateMany(context.Background(), q, bson.M{"$set": v})
	return Err(err)
}

// 查询条件和值
func UpdateByQMap(collection *mongo.Collection, q interface{}, v interface{}) bool {
	_, err := collection.UpdateMany(context.Background(), q, bson.M{"$set": v})
	return Err(err)
}

//------------------------

// 删除一条
func Delete(collection *mongo.Collection, q interface{}) bool {
	_, err := collection.DeleteOne(context.Background(), q)
	return Err(err)
}
func DeleteByIdAndUserId(collection *mongo.Collection, id, userId string) bool {
	_, err := collection.DeleteOne(context.Background(), GetIdAndUserIdQ(id, userId))
	return Err(err)
}
func DeleteByIdAndUserId2(collection *mongo.Collection, id, userId string) bool {
	_, err := collection.DeleteOne(context.Background(), GetIdAndUserIdBsonQ(id, userId))
	return Err(err)
}

// 删除所有
func DeleteAllByIdAndUserId(collection *mongo.Collection, id, userId string) bool {
	_, err := collection.DeleteMany(context.Background(), GetIdAndUserIdQ(id, userId))
	return Err(err)
}
func DeleteAllByIdAndUserId2(collection *mongo.Collection, id, userId string) bool {
	_, err := collection.DeleteMany(context.Background(), GetIdAndUserIdBsonQ(id, userId))
	return Err(err)
}

func DeleteAll(collection *mongo.Collection, q interface{}) bool {
	_, err := collection.DeleteMany(context.Background(), q)
	return Err(err)
}

func DropIndex(collection *mongo.Collection, indexName string) bool {
	_, err := collection.Indexes().DropOne(context.Background(), indexName)
	return Err(err)
}

//-------------------------

func Get(collection *mongo.Collection, id string, i interface{}) {
	collection.FindOne(context.Background(), bson.M{"_id": ObjectIDFromHex(id)}).Decode(i)
}
func Get2(collection *mongo.Collection, id string, i interface{}) {
	collection.FindOne(context.Background(), bson.M{"_id": ObjectIDFromHex(id)}).Decode(i)
}

func GetByQ(collection *mongo.Collection, q interface{}, i interface{}) {
	collection.FindOne(context.Background(), q).Decode(i)
}
func ListByQ(collection *mongo.Collection, q interface{}, i interface{}) {

	cur, err := collection.Find(context.Background(), q)
	if err != nil {
		log.Println(err)
		return
	}
	if err = cur.All(context.Background(), i); err != nil {
		log.Println(err)
		return
	}
}

func ListByQLimit(collection *mongo.Collection, q interface{}, i interface{}, limit int) {

	cursor, err := collection.Find(context.Background(), q, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return
	}
	// TODO 参数类型？
	if err = cursor.All(context.Background(), i); err != nil {
		return
	}

}

// 查询某些字段, q是查询条件, fields是字段名列表
func GetByQWithFields(collection *mongo.Collection, q bson.M, fields []string, i interface{}) {
	selector := make(bson.M, len(fields))
	for _, field := range fields {
		selector[field] = 1
	}

	cursor, err := collection.Find(context.Background(), q, options.Find().SetProjection(selector))

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.Background(), i); err != nil {
		return
	}
}

// 查询某些字段, q是查询条件, fields是字段名列表
func ListByQWithFields(collection *mongo.Collection, q bson.M, fields []string, i interface{}) {
	selector := make(bson.M, len(fields))
	for _, field := range fields {
		selector[field] = true
	}

	cursor, err := collection.Find(context.Background(), q, options.Find().SetProjection(selector))

	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.Background(), i); err != nil {
		return
	}
}
func GetByIdAndUserId(collection *mongo.Collection, id, userId string, i interface{}) {
	collection.FindOne(context.Background(), GetIdAndUserIdQ(id, userId)).Decode(i)
}
func GetByIdAndUserId2(collection *mongo.Collection, id, userId string, i interface{}) {
	collection.FindOne(context.Background(), GetIdAndUserIdBsonQ(id, userId)).Decode(i)
}

func AggregateQuery(collection *mongo.Collection, pipline mongo.Pipeline, i interface{}) {
	cursor, err := collection.Aggregate(context.Background(), pipline)
	if err != nil {
		log.Println(err)
		return
	}

	if err = cursor.All(context.Background(), i); err != nil {
		return
	}
}

// 按field去重
func Distinct(collection *mongo.Collection, q bson.M, field string, i interface{}) {
	// TODO _id
	pipline := mongo.Pipeline{
		{
			{"$match", q},
		},
		{
			{"$group", bson.D{{"_id", nil}, {field, bson.M{"$addToSet": "$" + field}}}},
		},
		{
			{"$project", bson.D{{"_id", 0}}},
		},
	}
	AggregateQuery(collection, pipline, i)
}

//----------------------

func Count(collection *mongo.Collection, q interface{}) int {
	ct, err := collection.CountDocuments(context.Background(), q)
	if err != nil {
		Err(err)
	}
	return int(ct)
}

func Has(collection *mongo.Collection, q interface{}) bool {
	return Count(collection, q) > 0
}

//-----------------

// 得到主键和userId的复合查询条件
func GetIdAndUserIdQ(id, userId string) bson.M {
	return bson.M{"_id": ObjectIDFromHex(id), "UserId": ObjectIDFromHex(userId)}
}
func GetIdAndUserIdBsonQ(id, userId string) bson.M {
	return bson.M{"_id": ObjectIDFromHex(id), "UserId": ObjectIDFromHex(userId)}
}

// DB处理错误
func Err(err error) bool {
	if err != nil {
		fmt.Println(err)
		// 删除时, 查找
		return err.Error() == "not found"
	}
	return true
}

// 检查mognodb是否lost connection
// 每个请求之前都要检查!!
func CheckMongoSessionLost() {
	// fmt.Println("检查CheckMongoSessionLostErr")
	err := client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("Lost connection to db!")
		// TODO 重连
	}
}
