package domain

type ListUsersOpts struct {
    ID       string
    FullName string
    Email    string
    Phone    string
    Role     string
    Status   string
    Page     int
    Limit    int
}