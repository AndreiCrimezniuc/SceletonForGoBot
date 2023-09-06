function getWatcher() { // now it is hardcode toDo: make axios to REST
    return [{
        "name": "Andrei Crimezniuc",
        "login": "@Andrei_crimezniuc                                                                                  "
    },{
        "name": "Георгий Сахечидзе",
        "login": "@ГеоргийДобавьUsername                                                                              "
    },{
        "name": "Nick Belousov",
        "login": "@tokimedo                                                                                           "
    },{
        "name": "Max Manole",
        "login": "@Max_Manole                                                                                         "
    },{
        "name": "Rodionov Igor",
        "login": "@bellylollipop                                                                                      "
    },{
        "name": "Constantin Cutureanu",
        "login": "@constantincutureanu                                                                                "
    },{
        "name": "Даниил Масленников",
        "login": "@dantes024                                                                                          "
    },{
        "name": "Игорь Кавлюк",
        "login": "@cavliman                                                                                           "
    },{
        "name": "Ярослав Флянку",
        "login": "@your_pixel                                                                                         "
    },{
        "name": "Росин Русановский",
        "login": "@rosin_rusanovschi                                                                                  "
    },]
}

function setWatchersForMonth(watchers, monthID) {
    const month = document.querySelector(`[key=${monthID}]`);
    const formMonth = document.createElement("form");
    
    // first micro
    // second micro
    // equipment

    watchers.map((w) => {
        const div = document.createElement("div");
        const node = document.createTextNode(w.name);
    })


}