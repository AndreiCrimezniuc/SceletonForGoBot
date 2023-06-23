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

function setWatchers(watchers) {
    const m1 = document.querySelector("[key=m1]");
    const m2 = document.querySelector("[key=m2]");
    const m3 = document.querySelector("[key=m3]");
    const m4 = document.querySelector("[key=m4]");

    watchers.map((w) => {
        const div = document.createElement("div");
        const node = document.createTextNode(w.name);

    })


}