import React, { useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, TableSortLabel } from '@material-ui/core';
import { ProjectDTO } from '../api/apis';

interface Props {
  projects: ProjectDTO[];
  setSelectedProject: React.Dispatch<React.SetStateAction<string | undefined>>;
}

const useStyles = makeStyles({
  table: {
    minWidth: 650,
    cursor: 'pointer'
  }
});

export default function ProjectTable({ setSelectedProject, projects }: Props) {
  const classes = useStyles();
  const [isDescending, setIsDescending] = useState(true);

  const handleCellClick = (selectedProject: string) => {
    setSelectedProject(selectedProject);
  };

  const sortByName = () => {
    setIsDescending(currentState => !currentState);

    function getComperator(a: string, b: string) {
      if (a > b) {
        return 1;
      }
      if (a < b) {
        return -1;
      }
      return 0;
    }

    projects.sort((a, b) => {
      if (isDescending) {
        return -getComperator(a.name, b.name);
      }
      return getComperator(a.name, b.name);
    });
  };

  return (
    <TableContainer>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>
              <TableSortLabel active direction={isDescending ? 'desc' : 'asc'} onClick={sortByName}>
                Name
              </TableSortLabel>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {projects.map((project, i) => (
            <TableRow hover key={i}>
              <TableCell component="th" scope="row" onClick={() => handleCellClick(project.id)}>
                {project.name}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
