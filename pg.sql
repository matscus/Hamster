create table tRole (id SERIAL,name varchar(40)UNIQUE,PRIMARY KEY (id, name))
insert into tRole (name)values('admin'),('user'),('admiral'),('inthebus')
create table tUsers (id SERIAL PRIMARY key,users varchar(40) UNIQUE,password varchar(180),password_Expiration TIMESTAMP,role varchar(40) references tRole(name))
insert into tUsers(users,password,password_expiration,role)values('god','f4e98344541784f2eabcf6fcd1daf050afd9a1bfa2c59819356fe0543752f311',to_timestamp(0),'admin')
create table tProjects (id SERIAL,name varchar(80) UNIQUE ,status varchar(40),PRIMARY KEY (id))
insert into tProjects (name,status)values('esfl','active'),('tarantool','active'),('telekard','active'),('afs','active'),('mdm','active'),('rkk','active'),('home_credit','active'),('pre_apruve_credits','active')
create table tUserProjects(id SERIAL PRIMARY key,user_id integer references tUsers(id) ON DELETE CASCADE,project_id integer references tProjects(id)ON DELETE CASCADE)
insert into tUserProjects (user_id,project_id)values(1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8)
create table tHosts (id SERIAL,ip varchar(40)UNIQUE,host_type varchar(40),users varchar(40),PRIMARY key(id));
insert into  tHosts (ip,host_type,users)values('127.0.0.1','monitoring','matscus')
create table tServices (id SERIAL,name varchar(80),host varchar(20)references tHosts(ip)ON DELETE CASCADE,uri varchar(80),type varchar(20),runstr varchar(256),PRIMARY KEY (id))
insert into tServices(name,host,uri,type,runstr)values('prometheus','127.0.0.1','http://localhost:9090/','prometheus','nohup ~/Hamster/bins/prometheus/prometheus --web.listen-address=localhost:9090 --config.file=./Hamster/bins/prometheus/prometheus.yml &> /dev/null')
create table tHostProjects(id SERIAL PRIMARY key,host_id integer references tHosts(id) ON DELETE CASCADE,project_id integer references tProjects(id)ON DELETE CASCADE)
create table tServiceProjects(id SERIAL PRIMARY key,service_id integer references tServices(id) ON DELETE CASCADE,project_id integer references tProjects(id)ON DELETE CASCADE)
insert into tServiceProjects (service_id,project_id)values(1,1),(1,2),(1,3),(1,4),(1,5),(1,6),(1,7),(1,8)
create table tScenarios (id SERIAL PRIMARY key,name varchar(40),test_type varchar(40),last_modified TIMESTAMP, gun_type varchar(40),project_name varchar(80) references tProjects(name),params jsonb)
create table tRuns (id SERIAL PRIMARY key,test_name varchar(40),test_type varchar(40),start_time TIMESTAMP,stop_time TIMESTAMP,status varchar(10),comment varchar(280),state varchar(10),project_name varchar(80) references tProjects(name))
CREATE OR REPLACE FUNCTION tUserProjects_ins_admin_project_function() RETURNS trigger AS 'BEGIN IF tg_op = ''INSERT'' THEN INSERT INTO tUserProjects(user_id,project_id) VALUES (1, new.id); RETURN new; END IF; END' LANGUAGE plpgsql;
CREATE TRIGGER new_tProjects_tg AFTER INSERT ON tProjects FOR each ROW EXECUTE PROCEDURE tUserProjects_ins_admin_project_function();
CREATE OR REPLACE FUNCTION tUserProjects_inc_function(integer,integer[]) RETURNS VOID AS $$ INSERT INTO tUserProjects(user_id, project_id) SELECT $1,i FROM unnest($2) i; $$ LANGUAGE sql STRICT;
create OR replace function tUsers_ins_function(u varchar(40),p varchar(180),r varchar(20)) RETURNS  integer AS $$ DECLARE id_val int; BEGIN insert into tUsers (users,password,password_expiration,role) values(u,p,now(),r) RETURNING id into id_val; RETURN id_val; END $$  LANGUAGE plpgsql;
CREATE OR replace function new_user_function(u varchar(40),p varchar(40),r varchar(40),pr integer[]) returns void  AS $$ DECLARE id_val int; begin select tUsers_ins_function(u,p,r) into id_val; PERFORM tUserProjects_inc_function(id_val,pr); end $$ LANGUAGE plpgsql;
create OR replace function tHosts_ins_function(u varchar(40),p varchar(180),r varchar(20)) RETURNS  integer AS $$ DECLARE id_val int; BEGIN insert into tHosts (ip,host_type,users) values (u,p,r) RETURNING id into id_val; RETURN id_val; END $$  LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION tHostProjects_inc_function(integer,integer[]) RETURNS VOID AS $$ INSERT INTO tHostProjects(host_id, project_id) SELECT  $1,i FROM unnest($2) i; $$ LANGUAGE sql STRICT;
CREATE OR replace function new_host_function(u varchar(40),p varchar(180),r varchar(20),pr integer[]) returns void  AS $$ DECLARE id_val int; begin select tHosts_ins_function(u,p,r) into id_val; PERFORM tHostProjects_inc_function(id_val,pr); end $$ LANGUAGE plpgsql;
create OR replace function tServices_ins_function(n varchar(80),h varchar(20),u varchar(180),t varchar(20),rstr varchar(256)) RETURNS  integer AS $$ DECLARE id_val int; BEGIN insert into tServices (name,host,uri,type,runstr) values (n,h,u,t,rstr) RETURNING id into id_val; RETURN id_val; END $$  LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION tServiceProjects_inc_function(integer,integer[]) RETURNS VOID AS $$ INSERT INTO tServiceProjects(service_id, project_id) SELECT  $1,i FROM unnest($2) i; $$ LANGUAGE sql STRICT;
CREATE OR replace function new_service_function(n varchar(80),h varchar(20),u varchar(180),t varchar(20),rstr varchar(256),pr integer[]) returns void  AS $$ DECLARE id_val int; begin select tServices_ins_function(n,h,u,t,rstr)  into id_val; PERFORM tServiceProjects_inc_function(id_val,pr); end $$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION tUserProjects_ins_user_function(_arr1 integer[], _arr2 integer[]) RETURNS VOID AS $$ INSERT INTO tUserProjects(user_id, project_id) SELECT * FROM unnest(_arr1, _arr2); $$ LANGUAGE sql STRICT;