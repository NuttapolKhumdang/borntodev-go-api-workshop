baseURL = "http://localhost:8080/workers"

try {
    fetch(baseURL)
        .then(r => r.json())
        .then(j => display(j))
        .catch(e => display([]))
} catch (e) {
    display([])
}

function display(data = []) {
    const target = document.querySelector("div");

    for (const k of data) {
        target.innerHTML += `<p>${k.ID}; ${k.Name} (${k.Email}, ${k.Phone}); Salary: THB ${k.Salary}</p>`
    }
}