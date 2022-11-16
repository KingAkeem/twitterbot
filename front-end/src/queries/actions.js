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

export async function getFollowing(username) {
  try {
    const data = await fetch(`http://localhost:8081/following/${username}`);
    const result = await data.json();
    return result;
  } catch(err) {
    console.error("Unable to fetch following. Error:", err)
  }
}

export async function getTweets(username) {
  try {
    const data = await fetch(`http://localhost:8081/tweets/${username}`);
    const result = await data.json();
    return result;
  } catch(err) {
    console.error("Unable to fetch followers. Error:", err)
  }
}