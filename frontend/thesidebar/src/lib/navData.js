import HomeIcon from '@mui/icons-material/Home';
import AssignmentIcon from '@mui/icons-material/Assignment';
import ArchiveIcon from '@mui/icons-material/Archive';
import SettingsIcon from '@mui/icons-material/Settings';
export const navData = [
    {
        id: 0,
        icon: <HomeIcon/>,
        text: "Home",
        link: "/"
    },
    {
        id: 1,
        icon: <AssignmentIcon/>,
        text: "Todos",
        link: "Todos"
    },
    {
        id: 2,
        icon: <ArchiveIcon/>,
        text: "Archive",
        link: "Archive"
    },
    {
        id: 3,
        icon: <SettingsIcon/>,
        text: "Settings",
        link: "settings"
    }
]