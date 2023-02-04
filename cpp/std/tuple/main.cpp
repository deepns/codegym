#include <iostream>
#include <array>
#include <tuple>

namespace
{
    std::tuple<std::string, double, int> nextStation(std::tuple<std::string, double, int> current_station)
    {
        // braced initialization of tuple supported from c++17 onwards
        return {
            std::get<0>(current_station),
            std::get<1>(current_station) + 1,
            std::get<2>(current_station)};
    }

    typedef std::tuple<int, std::string> DayOfWeek;
    const std::array<std::string, 7> WeekDays = {
        "Sunday",
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday",
        "Saturday"};

    DayOfWeek nextDay(DayOfWeek day)
    {
        auto next_day_index = (std::get<0>(day) + 1 % WeekDays.size());
        return {
            next_day_index,
            WeekDays[next_day_index]};
    }

} // namespace

int main()
{
    // pack arbitrary number of values
    std::tuple<std::string, double, int> station = {"npr", 91.5, 24060};

    // how to read individual value from a tuple?
    // can't read directly. need to use the template function std::get
    std::cout << "station<0>=" << std::get<0>(station)
              << " station<1>=" << std::get<1>(station)
              << " station<2>=" << std::get<2>(station)
              << std::endl;

    // unpacking values from a tuple
    // using std::tie on the lhs

    // if the type has already been declared
    std::string name;
    double freq;
    int zip;

    std::tie(name, freq, zip) = station;
    std::cout << "name=" << name
              << " freq=" << freq
              << " zip=" << zip
              << std::endl;

    // if the types haven't been declared already,
    // can use structural bindings (from C++17)
    // https://riptutorial.com/cplusplus/example/3384/structured-bindings
    auto [s_name, s_freq, s_zip] = nextStation(station);
    std::cout << "s_name=" << s_name
              << " s_freq=" << s_freq
              << " s_zip=" << s_zip
              << std::endl;

    // another example
    DayOfWeek sunday = DayOfWeek{0, WeekDays[0]};
    auto monday = nextDay(sunday);

    // can selectively ignore values during unpacking using std::ignore
    std::string day;
    std::tie(std::ignore, day) = nextDay(monday);
    std::cout << "day=" << day
              << std::endl;

    // can make tuple with std::make_tuple as well
    auto codes = std::make_tuple(101, 95.6, 97.8, 102);
    std::cout << "codes<0>=" << std::get<0>(codes) << std::endl;
}