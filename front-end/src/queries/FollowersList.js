import * as React from 'react';

import {
	Avatar,
	Divider,
	List,
	ListItem,
	ListItemAvatar,
	ListItemText
} from '@mui/material';

import { useParams } from 'react-router-dom';
import { getFollowers } from './actions';

const list = (user) => {
	const fields = ["username", "id", "name", "description", "location", "url"];
	return fields.map(
		(field, index) => {
			return (
				<>
					<ListItem alignItems="flex-start">
						{field === "username" ? 
							<ListItemAvatar>
								<Avatar
									alt={user[field]}
									src={user["profile_image_url"]}
								/>
							</ListItemAvatar> : null}
						<ListItemText
							primary={field.toUpperCase()}
							primaryTypographyProps={{color: "primary"}}
							secondary={
							<React.Fragment>
								{user[field]}
							</React.Fragment>
							}
						/>
					</ListItem>
				</>
			);
		}
	)
}


export default function FollowList(props) {
	const username = useParams()["username"];

	const [followList, setFollowList] = React.useState([]);
	React.useEffect(() => {
		getFollowers(username).then(data => setFollowList(data));
	}, [username]);

	return (
		<List sx={{ width: '100%', maxWidth: 360, bgcolor: 'background.paper' }}>
			{followList.map(user => {
				return (
					<>
					{list(user)}
					<Divider/>
					</>
				);
			})}
		</List>
	)
}