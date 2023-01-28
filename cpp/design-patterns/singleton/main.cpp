// Design patterns - Singleton
//
// When to use?
//  1. Need single point of access to a shared resource (e.g. database, cache
//     file etc.)
//  2. Ensure the single point of access by restricting the class to have only
//     one object
//
// A way to use?
//  - Make the default constructor private, so other objects can't instantiate
//    the singleton from within the class
//  - Define a static method that acts as the constructor. Construct the class
//    from the static method.
//
// Notes
//  - Protect the instantiation with a lock to prevent creation of multiple
//    instances when called from multiple threads

#include <memory>
#include <iostream>
#include <map>

class DBConnection {
public:
    
    // Should not be cloneable
    DBConnection(DBConnection &other) = delete;
    
    // Should not be assignable
    void operator=(const DBConnection &) = delete;
        
    static std::shared_ptr<DBConnection> GetInstance();

    std::string Get(const std::string& key) {
        return db_[key];
    }

    void Put(const std::string& key, const std::string& val) {
        db_[key] = val;
    }

private:
    static std::shared_ptr<DBConnection> instance_;
    std::map<std::string, std::string> db_;
};

std::shared_ptr<DBConnection> DBConnection::GetInstance() {
        if (instance_ == nullptr) {
            instance_ = std::make_shared<DBConnection>();
        }
        return instance_;
    }
    
int main() {
    std::shared_ptr<DBConnection> db = DBConnection::GetInstance();

    db->Put("foo", "bar");
    std::cout << "foo=" << db->Get("foo") << std::endl;

    return 0;
}