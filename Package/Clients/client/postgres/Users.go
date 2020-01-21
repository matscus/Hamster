package postgres

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
	pg "github.com/lib/pq"
)

//User struct data user for returln all users
type User struct {
	ID       int64    `json:"id,omitempty"`
	User     string   `json:"user,omitempty"`
	Password string   `json:"password,omitempty"`
	Role     string   `json:"role,omitempty"`
	Projects []string `json:"projects,omitempty"`
}

//NewUser - func create new users
func (c PGClient) NewUser(users string, password string, role string, projects []string) error {
	_, err := c.DB.Query("select new_user_function($1,$2,$3,$4) ", users, password, role, pg.Array(projects))
	return err
}

//UpdateUser - func for update  users, if len slice project ==0, projects not updatet
func (c PGClient) UpdateUser(id string, role string, projects []string) error {
	_, err := c.DB.Exec("update tUsers set role=$2 where id=$3", role, id)
	if err != nil {
		return err
	}
	if len(projects) != 0 {
		_, err := c.DB.Exec("delete from tServiceProjects where user_id=$1", id)
		if err != nil {
			return err
		}
		_, err = c.DB.Query("select tServiceProjects_inc_function($1,$2)", id, pg.Array(projects))
		if err != nil {
			return err
		}
	}
	return err
}

//DeleteUser - delete user
func (c PGClient) DeleteUser(id int64) (err error) {
	_, err = c.DB.Exec("delete from tUsers where id=$1", id)
	return err
}

//GetUserHash - return user password hash
func (c PGClient) GetUserHash(user string) (hash string, err error) {
	err = c.DB.QueryRow("select password from tUsers where users=$1", user).Scan(&hash)
	return hash, err
}

//ChangeUserPassword - delete user
func (c PGClient) ChangeUserPassword(id int64, password string) (err error) {
	t := time.Now().Add(2880 * time.Hour).Unix()
	_, err = c.DB.Exec("update tUsers set password=$1, password_expiration=to_timestamp($2) where id=$3", password, t, id)
	return nil
}

//GetUserPasswordExp - return user password expiration
func (c PGClient) GetUserPasswordExp(user string) (exp string, err error) {
	err = c.DB.QueryRow("select password_expiration from tUsers where users=$1", user).Scan(&exp)
	return exp, err
}

//GetAllUsers - return user password expiration
func (c PGClient) GetAllUsers() (users []User, err error) {
	var rows *sql.Rows
	rows, err = c.DB.Query("select id,users,role from tUsers")
	if err != nil {
		return nil, err
	}
	res := make([]User, 0, 20)
	for rows.Next() {
		u := User{}
		rows.Scan(&u.ID, &u.User, &u.Role)
		rowsProjects, err := c.DB.Query("select name from tProjects where id in(select project_id from tuserprojects where user_id=$1)", u.ID)
		if err != nil {
			return nil, err
		}
		for rowsProjects.Next() {
			var project string
			err = rowsProjects.Scan(&project)
			if err != nil {
				return nil, err
			}
			u.Projects = append(u.Projects, project)
		}
		u.Password = "so that the password is not stored in clear text, use https(с) Programmer of Mail.ru Group"
		res = append(res, u)
	}
	return res, nil
}

//GetUserRoleAndProjects - func to return all users role and  project for front project and role models
func (c PGClient) GetUserRoleAndProjects(user string) (role string, projects []string, err error) {
	err = c.DB.QueryRow("select role,projects from tUsers where users=$1", user).Scan(&role, pq.Array(&projects))
	return role, projects, err
}

//GetUserIDAndRole - func to return all users role and  project for front project and role models
func (c PGClient) GetUserIDAndRole(user string) (id int64, role string, err error) {
	if err := c.DB.QueryRow("select id,role from tUsers where users=$1", user).Scan(&id, &role); err != nil {
		return 0, role, err
	}
	return id, role, nil
}

//GetUserProjects - func to return all users role and  project for front project and role models
func (c PGClient) GetUserProjects(userID int64) (projects []string, err error) {
	rows, err := c.DB.Query("select name from tProjects where id in(select project_id from tuserprojects where user_id=$1)", userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var project string
		err = rows.Scan(&project)
		if err != nil {
			return nil, err
		} else {
			projects = append(projects, project)
		}
	}
	return projects, nil
}

//UpdatetUserProjects -
func (c PGClient) UpdatetUserProjects(id int64, projects []string) error {
	_, err := c.DB.Exec("delete from tUserProjects where user_id=$1", id)
	if err != nil {
		return err
	}
	_, err = c.DB.Query("select tUserProjects_inc_function($1,$2)", id, pg.Array(projects))
	if err != nil {
		return err
	}
	return nil
}
