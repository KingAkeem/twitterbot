export async function getUser(username) {
  try {
    const data = await fetch(`http://localhost:8081/query?username=${username}`);
    const result = await data.json();
    return result;
  } catch(err) {
    console.error("Unable to fetch user. Error:", err)
  }
}