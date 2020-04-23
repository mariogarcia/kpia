-- tables
CREATE TABLE piweek (
    id UUID PRIMARY KEY, 
    year integer, 
    month integer, 
    begins timestamp, 
    ends timestamp
); 

CREATE TABLE projects (
    id UUID PRIMARY KEY, 
    name text, 
    description text, 
    piweek_id UUID REFERENCES piweek (id)
);

CREATE TABLE members (
    id UUID PRIMARY KEY, 
    name text, 
    profile text
);

CREATE TABLE projects_members (
    project_id UUID REFERENCES projects (id), 
    member_id UUID REFERENCES members (id)
);
