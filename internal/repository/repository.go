package repository
import (
	"database/sql"
	"fmt"
	"github.com/getclasslabs/user/internal/config"
	"github.com/getclasslabs/user/tools"
	_ "github.com/go-sql-driver/mysql"
)
type Repo interface {
	Connect()
	Insert(*tools.Infos, string, ...interface{}) (sql.Result, error)
	Update(*tools.Infos, string, ...interface{}) (sql.Result, error)
	Get(*tools.Infos, string, ...interface{}) (map[string]interface{}, error)
	Fetch(*tools.Infos, string, ...interface{}) ([]map[string]interface{}, error)
}

type MySQL struct {
	db *sql.DB
}

func (m MySQL) Connect() {
	c := config.Config.Mysql
	connectionLine := "%s:%s@tcp(%s:%d)/%s"
	connectionLine = fmt.Sprintf(connectionLine,
		c.User, c.Password, c.Host, c.Port, c.Database)
	drive, err := sql.Open("mysql", connectionLine)
	if err != nil {
		panic(err)
	}
	m.db = drive
}

func (m MySQL) Insert(i *tools.Infos, query string, args ...interface{}) (sql.Result, error) {
	i.Span = tools.TraceIt(i, "inserting")
	defer i.Span.Finish()

	var result sql.Result

	stmtIns, err := m.db.Prepare(query)
	if err != nil {
		i.LogError(err)
		return result, err
	}
	defer stmtIns.Close()

	result, err = stmtIns.Exec(args...)
	if err != nil {
		i.LogError(err)
		return result, err
	}

	return result, nil
}

func (m MySQL) Update(i *tools.Infos, query string, args ...interface{}) (sql.Result, error) {
	i.Span = tools.TraceIt(i, "updating")
	defer i.Span.Finish()

	var result sql.Result

	stmtIns, err := m.db.Prepare(query)
	if err != nil {
		i.LogError(err)
		return result, err
	}
	defer stmtIns.Close()

	result, err = stmtIns.Exec(args...)
	if err != nil {
		i.LogError(err)
		return result, err
	}

	return result, nil
}

func (m MySQL) Get(i *tools.Infos, query string, args ...interface{}) (map[string]interface{}, error) {
	i.Span = tools.TraceIt(i, "select one")
	defer i.Span.Finish()

	stmt, err := m.db.Prepare(query)
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	entry := make(map[string]interface{})
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry = make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
	}
	return entry, nil
}

func (m MySQL) Fetch(i *tools.Infos, query string, args ...interface{}) ([]map[string]interface{}, error) {
	i.Span = tools.TraceIt(i, "select many")
	defer i.Span.Finish()

	stmt, err := m.db.Prepare(query)
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		i.LogError(err)
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}


var Repository Repo = MySQL{}