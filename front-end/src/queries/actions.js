export async function getUser(username) {
  try {
    const data = await fetch(`http://localhost:8081/user/${username}`);
    const result = await data.json();
    return result;
  } catch(err) {
    console.error("Unable to fetch user. Error:", err)
  }
}

export async function getFollowers(username) {
  try {
    const data = await fetch(`http://localhost:8081/followers/${username}`);
    const result = await data.json();
    return result;
  } catch(err) {
    console.error("Unable to fetch followers. Error:", err)
  }
}