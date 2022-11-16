import * as React from 'react';

import {
	Divider,
	List,
	ListItemText,
	ListItem,
} from '@mui/material';

import { useParams } from 'react-router-dom';
import { getTweets } from './actions';

const list = (user) => {
	const fields = ["id", "text", "lang", "source"];
	return fields.map(
		(field, index) => {
			return (
				<>
					<ListItem alignItems="flex-start">
						<ListItemText
							primary={field[0].toUpperCase() + field.substring(1)}
							primaryTypographyProps={{color: "primary"}}
							secondary={
							<React.Fragment>
								{user[field]}
							</React.Fragment>
							}
						/>
					</ListItem>
					<Divider variant={field === "username" ? "inset" : "fullWidth"} component="li" /> 
				</>
			);
		}
	)
}


export default function TweetList(props) {
	const username = useParams()["username"];

	const [tweetList, setTweetList] = React.useState([]);
	React.useEffect(() => {
		getTweets(username).then(data => setTweetList(data));
	}, [username]);

	return (
		<List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>
			{tweetList && tweetList.map(tweet => list(tweet))}
		</List>
	)
}