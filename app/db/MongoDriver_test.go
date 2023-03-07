package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"leanote/app/info"
	"log"
	"reflect"
	"sort"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGet(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		id         string
		i          interface{}
	}
	var user bson.D
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				"5368c1aa99c37b029d000001",
				&user,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get(tt.args.collection, tt.args.id, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestGet2(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		id         string
		i          interface{}
	}
	var user bson.D
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				"5368c1aa99c37b029d000001",
				&user,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Get2(tt.args.collection, tt.args.id, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestGetByQ(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          interface{}
		i          interface{}
	}
	userNotebooks := []info.Notebook{}
	orQ := []bson.M{
		//IsDeleted为false 或者 IsDeleted值不存在
		bson.M{"IsDeleted": false},
		bson.M{"IsDeleted": bson.M{"$exists": false}},
	}
	query := bson.M{"UserId": ObjectIDFromHex("540817e099c37b583c000001"), "$or": orQ}
	//query := bson.M{"UserId": ObjectIDFromHex("540817e099c37b583c000001")}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

		{"nm",
			args{
				Notebooks,
				query,
				//&bson.D{},
				&userNotebooks,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//GetByQ(tt.args.collection, tt.args.q, tt.args.i)
			ListByQ(tt.args.collection, tt.args.q, tt.args.i)
			log.Printf("【%v】\n", reflect.TypeOf(userNotebooks))
			log.Printf("【%v】\n", len(userNotebooks))
			log.Printf("【%v】\n", userNotebooks)

			a := ParseAndSortNotebooks(
				userNotebooks,
				true,
				true,
			)
			log.Printf("【%v】\n", reflect.TypeOf(a))
			log.Printf("【%v】\n", len(a))
			log.Printf("【%v】\n", a)
		})
	}
}

func TestListByQ(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          interface{}
		i          interface{}
	}

	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				// bson.M{"Username": "admin"},
				bson.M{},
				&bson.A{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListByQ(tt.args.collection, tt.args.q, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestListByQLimit(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          interface{}
		i          interface{}
		limit      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				bson.M{"Username": "admin"},
				&bson.A{},
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListByQLimit(tt.args.collection, tt.args.q, tt.args.i, tt.args.limit)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestGetByQWithFields(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          bson.M
		fields     []string
		i          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				bson.M{"Username": "admin"},
				[]string{"Username"},
				&bson.A{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetByQWithFields(tt.args.collection, tt.args.q, tt.args.fields, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestListByQWithFields(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          bson.M
		fields     []string
		i          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				bson.M{"Username": "admin"},
				[]string{"Username"},
				&bson.A{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListByQWithFields(tt.args.collection, tt.args.q, tt.args.fields, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestGetByIdAndUserId(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		id         string
		userId     string
		i          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Notes,
				"540817e099c37b583c000005",
				"540817e099c37b583c000001",
				&bson.D{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetByIdAndUserId(tt.args.collection, tt.args.id, tt.args.userId, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestGetByIdAndUserId2(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		id         string
		userId     string
		i          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Notes,
				"540817e099c37b583c000005",
				"540817e099c37b583c000001",
				&bson.D{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetByIdAndUserId2(tt.args.collection, tt.args.id, tt.args.userId, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestDistinct(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          bson.M
		field      string
		i          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				bson.M{"Username": "admin"},
				// bson.M{},
				"Username",
				&bson.A{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Distinct(tt.args.collection, tt.args.q, tt.args.field, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				bson.M{"Username": "admin"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Count(tt.args.collection, tt.args.q)
			log.Printf("【%v】\n", got)
		})
	}
}

func TestHas(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		q          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				bson.M{"Username": "admin"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Has(tt.args.collection, tt.args.q)
			log.Printf("【%v】\n", got)
		})
	}
}

func TestAggregateQuery(t *testing.T) {
	type args struct {
		collection *mongo.Collection
		pipline    mongo.Pipeline
		i          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"nm",
			args{
				Users,
				// bson.M{"Username": "admin"},
				mongo.Pipeline{
					{
						{"$match", bson.M{}},
					},
					{
						{"$project", bson.M{"Username": 1, "_id": 0}},
					},
					{
						{"$sort", bson.D{{"Username", 1}}},
					},
				},
				&bson.A{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AggregateQuery(tt.args.collection, tt.args.pipline, tt.args.i)
			log.Printf("【%v】\n", tt.args.i)
		})
	}
}

func ParseAndSortNotebooks(userNotebooks []info.Notebook, noParentDelete, needSort bool) info.SubNotebooks {
	// 整理成info.Notebooks
	// 第一遍, 建map
	// notebookId => info.Notebooks
	userNotebooksMap := make(map[primitive.ObjectID]*info.Notebooks, len(userNotebooks))
	for _, each := range userNotebooks {
		newNotebooks := info.Notebooks{Subs: info.SubNotebooks{}}
		newNotebooks.NotebookId = each.NotebookId
		newNotebooks.Title = each.Title
		//		newNotebooks.Title = html.EscapeString(each.Title)
		newNotebooks.Title = strings.Replace(strings.Replace(each.Title, "<script>", "", -1), "</script", "", -1)
		newNotebooks.Seq = each.Seq
		newNotebooks.UserId = each.UserId
		newNotebooks.ParentNotebookId = each.ParentNotebookId
		newNotebooks.NumberNotes = each.NumberNotes
		newNotebooks.IsTrash = each.IsTrash
		newNotebooks.IsBlog = each.IsBlog

		// 存地址
		userNotebooksMap[each.NotebookId] = &newNotebooks
	}

	// 第二遍, 追加到父下

	// 需要删除的id
	needDeleteNotebookId := map[primitive.ObjectID]bool{}
	for id, each := range userNotebooksMap {
		// 如果有父, 那么追加到父下, 并剪掉当前, 那么最后就只有根的元素
		if !each.ParentNotebookId.IsZero() {
			if userNotebooksMap[each.ParentNotebookId] != nil {
				userNotebooksMap[each.ParentNotebookId].Subs = append(userNotebooksMap[each.ParentNotebookId].Subs, each) // Subs是存地址
				// 并剪掉
				// bug
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			} else if noParentDelete {
				// 没有父, 且设置了要删除
				needDeleteNotebookId[id] = true
				// delete(userNotebooksMap, id)
			}
		}
	}

	// 第三遍, 得到所有根
	final := make(info.SubNotebooks, len(userNotebooksMap)-len(needDeleteNotebookId))
	i := 0
	for id, each := range userNotebooksMap {
		if !needDeleteNotebookId[id] {
			final[i] = each
			i++
		}
	}

	// 最后排序
	if needSort {
		return sortSubNotebooks(final)
	}
	return final
}

// 排序
func sortSubNotebooks(eachNotebooks info.SubNotebooks) info.SubNotebooks {
	// 遍历子, 则子往上进行排序
	for _, eachNotebook := range eachNotebooks {
		if eachNotebook.Subs != nil && len(eachNotebook.Subs) > 0 {
			eachNotebook.Subs = sortSubNotebooks(eachNotebook.Subs)
		}
	}

	// 子排完了, 本层排
	sort.Sort(&eachNotebooks)
	return eachNotebooks
}

func Test_getUrl(t *testing.T) {
	type args struct {
		url    string
		dbname string
	}
	tests := []struct {
		name string
		args args
	}{
		{"nm",
			args{
				"", "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getUrl(tt.args.url, tt.args.dbname)
			log.Printf("【%v,%v】", got, got1)
		})
	}
}
