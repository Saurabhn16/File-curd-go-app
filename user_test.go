func TestCreateUser(t *testing.T) {
    // Set up a mock MongoDB database
    db := new(MockDatabase)
    db.On("CreateUser", mock.Anything).Return(nil)

    // Create a new UserAPI using the mock database
    api := UserAPI{db}

    // Create a new user
    user := User{
        Name: "Test User",
        Email: "test@example.com",
    }
    err := api.CreateUser(user)

    // Check that the CreateUser function returns nil (no error)
    if err != nil {
        t.Errorf("Expected CreateUser to return nil, but got %v", err)
    }

    // Check that the CreateUser function was called with the correct arguments
    db.AssertCalled(t, "CreateUser", user)
}